include .env

run-bot:
	RPC_URL=$(RPC_URL) PRIVATE_KEY=$(PRIVATE_KEY) go run .

scores:
	RPC_URL=$(RPC_URL) go run scripts/scores.go

contracts-bindings:
	abigen --abi ./contract/frontrunner/abi.json --pkg frontrunner --type Frontrunner --out ./contract/frontrunner/frontrunner.go
