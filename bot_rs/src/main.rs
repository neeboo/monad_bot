use alloy::primitives::Address;
use anyhow::{anyhow, Result};
use dotenv::dotenv;
use frontrunner_bot::FrontrunnerBot;
use std::env;
use std::str::FromStr;
use tracing::level_filters::LevelFilter;
use tracing_subscriber::EnvFilter;

#[tokio::main]
async fn main() -> Result<()> {
    dotenv().ok();
    init_logger()?;

    let rpc_url =
        env::var("RPC_URL").map_err(|e| anyhow!(format!("not found rpc_url {}", e.to_string())))?;
    let private_key = env::var("PRIVATE_KEY")
        .map_err(|e| anyhow!(format!("not found private_key {}", e.to_string())))?;
    let address =
        env::var("ADDRESS").map_err(|e| anyhow!(format!("not found address {}", e.to_string())))?;

    let bot =
        FrontrunnerBot::new(rpc_url, private_key, Address::from_str(address.as_str())?).await?;

    // loop run 100 times
    let max = 100;
    let mut count = 0;
    let _bot = bot.clone();
    while count < max {
        match _bot.run_bot().await {
            Ok(_) => {
                println!("success: #{}", count);
            }
            Err(_) => {
                println!("error: #{}", count);
            }
        };
        tokio::time::sleep(tokio::time::Duration::from_secs(1)).await;
        count += 1;
    }

    Ok(())
}

pub fn init_logger() -> Result<()> {
    let filter = EnvFilter::try_from_default_env().unwrap_or(
        EnvFilter::new("info")
            .add_directive(LevelFilter::ERROR.into())
            .add_directive(LevelFilter::WARN.into())
            .add_directive(LevelFilter::INFO.into()),
    );
    tracing_subscriber::fmt::fmt()
        .with_timer(tracing_subscriber::fmt::time::ChronoLocal::new(
            "[%Y-%m-%d %H:%M:%S%.3f]".into(),
        ))
        .with_env_filter(filter)
        .init();

    Ok(())
}
