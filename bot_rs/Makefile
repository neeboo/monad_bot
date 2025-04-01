include .env

run:
	cargo run --release

test:
	cargo test

contract-bindings:
	alloy bind --abi ../bot_go/contracts/frontrunner.abi.json --out-dir ./src/contracts
