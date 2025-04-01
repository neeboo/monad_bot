include .env

run:
	cargo run --release

test:
	cargo test

bench:
	cargo bench

contract-bindings:
	alloy bind --abi ./contracts/frontrunner.abi.json --out-dir ./src/contracts
