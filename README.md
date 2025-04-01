# Monad Bot
An experimental *Vibe Coding* project.
1. I forked [https://github.com/FastLane-Labs/break-monad-frontrunner-bot](https://github.com/FastLane-Labs/break-monad-frontrunner-bot) to `bot_go` folder
2. I tried using prompts and CLine+Deepseek(Reasoner as Planner, Chat as Actor) to rewrite the `bot_go` to `bot_rs`
3. After a few hours it finally lands however still need my engineer work a lot

## How to run
1. `cd bot_rs`
1. change `.env.example` to `.env` in bot_rs
2. replace the `RPC URL` to Monad Testnet, eg: "https://rpc.monad-testnet-2.fastlane.xyz/b3qFoDfY9sR44yRyOeHAfyj9dpEXVoOC"
3. replace the `PRIVATE_KEY` to your actuall private key, start with `0x`, *NOT SEEDPHRASE*
4. make sure you have some `MON` in your `Wallet`, which should be the address of the associated `PRIVATE_KEY`
5. then run `cargo run` or `cargo build --release && ./target/release/frontrunner-bot`

## How to build yourself?
1. change anything in `main.rs` or `lib.rs`
2. `cargo build` or `cargo run`

## Findings
1. It's hard to use `planner` just to say `hey, rewrite everything` then switch to `Act Mode` and boom! It doesn't work, at all.
2. LLM models are not familiar with latest software dependencies, you need to manually adjust those deps version if they bindly adding them.
3. For MVC development model, the `M` comes first, so you need to let LLM knows your model design, and let them draft the interface/trait/types, then start writing specific functions.
4. When it comes to fix things, you need to manually fix bugs when the LLM shows some speculating codes.


