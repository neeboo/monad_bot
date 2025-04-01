use alloy::network::{EthereumWallet, NetworkWallet, TransactionBuilder};
use alloy::primitives::{Address, Bytes, ChainId};
use alloy::providers::ProviderBuilder;
use alloy::rpc::types::TransactionReceipt;
use alloy::sol;
use alloy_provider::fillers::{
    BlobGasFiller, ChainIdFiller, FillProvider, GasFiller, JoinFill, NonceFiller,
};
use alloy_provider::{Identity, PendingTransactionError, Provider, RootProvider};
use alloy_signer::Signer;
use alloy_signer_local::PrivateKeySigner;
use anyhow::{Context, Result};
use std::str::FromStr;
use tracing::{error, info};

pub mod analyzer;
pub mod types;

use crate::Frontrunner::ParticipantData;
use analyzer::ArbitrageAnalyzer;
use types::{ArbitrageOpportunity, TransactionAnalysis};

pub type EvmProvider = FillProvider<
    JoinFill<
        Identity,
        JoinFill<GasFiller, JoinFill<BlobGasFiller, JoinFill<NonceFiller, ChainIdFiller>>>,
    >,
    RootProvider,
>;

sol!(
    #[allow(missing_docs)]
    #[sol(rpc)]
    Frontrunner,
    "contract/abi.json",
);

#[derive(Debug, Clone)]
pub struct FrontrunnerBot {
    provider: EvmProvider,
    wallet: EthereumWallet,
    contract_address: Address,
}

impl FrontrunnerBot {
    pub async fn new(
        rpc_url: String,
        private_key: String,
        contract_address: Address,
    ) -> Result<Self> {
        let mut signer = PrivateKeySigner::from_str(private_key.as_str())?;

        let provider = ProviderBuilder::new().on_http(rpc_url.parse()?);

        let chain_id = provider.get_chain_id().await?;
        signer.set_chain_id(Some(ChainId::from(chain_id)));
        let wallet = EthereumWallet::from(signer);

        Ok(Self {
            provider,
            wallet,
            contract_address,
        })
    }

    pub async fn run_bot(&self) -> Result<()> {
        let chain_id = self.provider.get_chain_id().await?;
        info!("Actual Chain ID: {}", chain_id);
        let (fee_cap, tip_cap) = self.get_gas_price().await?;
        let balance = self.get_balance().await?;
        info!(
            "Current gas price - maxFeePerGas: {}, maxPriorityFeePerGas: {}",
            fee_cap, tip_cap
        );
        info!("contract address {}", self.contract_address);
        info!(
            "Wallet Address {} Balance: {}",
            self.wallet.default_signer().address(),
            balance
        );

        let nonce = self
            .provider
            .get_transaction_count(self.wallet.default_signer().address())
            .await
            .context("Failed to get nonce")?;

        info!("nonce is {}", nonce);

        let contract = Frontrunner::new(self.contract_address, self.provider.clone());
        let mut builder = contract.frontrun();

        let request = builder.into_transaction_request();

        let gas_limit = self
            .provider
            .estimate_gas(&request)
            .await
            .context("Failed to estimate gas")?;

        info!("gas_limit is {}", gas_limit);

        let _tx = request
            .from(self.wallet.default_signer().address())
            .max_fee_per_gas(fee_cap)
            .max_priority_fee_per_gas(tip_cap)
            .with_gas_limit(gas_limit)
            .with_chain_id(ChainId::from(chain_id))
            .with_nonce(nonce)
            .build(&self.wallet)
            .await?;

        let provider = self.provider.clone();
        let tx_string = hex::encode(_tx.tx_hash().0);
        let tx_clone = _tx.clone();
        info!("pending tx is {}", tx_string);
        let mut pending = provider.send_tx_envelope(tx_clone).await?;
        // pending.set_required_confirmations(1);
        match pending.get_receipt().await {
            Ok(re) => {
                if re.status() == true {
                    info!(
                        "Confirmed frontrun: https://monad-testnet.socialscan.io/tx/{}",
                        tx_string
                    );
                } else {
                    error!(
                        "Frontrun reverted: https://monad-testnet.socialscan.io/tx/{}",
                        tx_string
                    );
                }
            }
            Err(_) => {
                error!(
                    "Frontrun reverted: https://monad-testnet.socialscan.io/tx/{}",
                    tx_string
                );
            }
        }
        Ok(())
    }

    pub async fn get_gas_price(&self) -> Result<(u128, u128)> {
        let fee_cap = self.provider.get_gas_price().await?;
        let tip_cap = self.provider.get_max_priority_fee_per_gas().await?;

        Ok((fee_cap, tip_cap))
    }

    pub async fn get_balance(&self) -> Result<String> {
        let balance = self
            .provider
            .get_balance(self.wallet.default_signer().address())
            .await?;
        Ok(balance.to_string())
    }

    pub async fn scores(&self) -> Result<Vec<ParticipantData>> {
        let (fee_cap, tip_cap) = self.get_gas_price().await?;

        info!(
            "Sending frontrun - maxFeePerGas: {}, maxPriorityFeePerGas: {}",
            fee_cap, tip_cap
        );

        let contract = Frontrunner::new(self.contract_address, self.provider.clone());

        let builder = contract
            .getScores()
            .from(self.wallet.default_signer().address())
            .max_fee_per_gas(fee_cap)
            .max_priority_fee_per_gas(tip_cap);

        let sc = builder.call().await?;
        Ok(sc._0)
    }
}
