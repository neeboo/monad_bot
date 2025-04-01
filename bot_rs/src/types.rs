use alloy::primitives::{Address, U256};
use alloy::rpc::types::eth::Transaction;

#[derive(Debug, Clone)]
pub struct ArbitrageOpportunity {
    pub target_tx: Transaction,
    pub profit_estimate: U256,
    pub gas_cost: U256,
    pub risk_factor: f64,
    pub deadline: u64,
}

#[derive(Debug, Clone)]
pub struct TransactionAnalysis {
    pub is_arbitrage: bool,
    pub exchange_address: Option<Address>,
    pub token_pair: Option<(Address, Address)>,
    pub price_difference: Option<U256>,
}
