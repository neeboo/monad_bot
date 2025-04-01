include .env
build:
	cargo build --release

run:
	cargo run --release

test:
	cargo test

contract-bindings:
	alloy bind --abi ../bot_go/contracts/frontrunner.abi.json --out-dir ./src/contracts
