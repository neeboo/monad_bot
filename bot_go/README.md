# FastLane Frontrunner Bot

💡We also have a python version of the bot [Here 🐍](https://github.com/FastLane-Labs/break-monad-frontrunner-bot-py)

<p align="center">
  <img src="frontrunner-gif.gif" alt="Frontrunner Game Animation" width="600">
</p>

The objective is simple, you compete as a searcher to be first to land a transaction on-chain in each new block.
Your ranking is determined by your win/loss ratio weighted by number of attempts.

## Prerequisites

### 1. Install Go (Golang)

First, you'll need to install Go on your computer:

#### Windows:

1. Download the installer from [Go's official website](https://golang.org/dl/)
2. Run the installer and follow the prompts
3. Open Command Prompt and verify installation:

```sh
go version
```

#### Mac:

Using Homebrew:

```sh
brew install go
```

#### Linux (Ubuntu/Debian):

```sh
sudo apt update
sudo apt install golang
```

### 2. Set up your environment

1. Create a `.env` file in the project root directory
2. Generate a private key (if you don't have one):

   ```sh
   # Using OpenSSL (recommended)
   openssl rand -hex 32

   # Alternative: Using Python
   python3 -c "import secrets; print(secrets.token_hex(32))"
   ```

3. To get your Ethereum address from the private key, use Python:

   ```sh
   # Install web3 if you haven't already
   pip install web3

   # Run this Python command (replace YOUR_PRIVATE_KEY with the key generated above)
   python3 -c "from web3 import Web3; w3 = Web3(); private_key = 'YOUR_PRIVATE_KEY'; account = w3.eth.account.from_key('0x' + private_key); print(f'Private key: {private_key}'); print(f'Public address: {account.address}')"
   ```

   Additionally, your private key can sometimes be found in your wallet settings.

4. Add your configuration:

```sh
PRIVATE_KEY=your_private_key_here
RPC_URL=your_rpc_url_here
```

⚠️ IMPORTANT SECURITY NOTES:

- NEVER share your private key or commit it to version control!
- Store your private key securely and keep a backup

Additional RPC URLs:

```
London: https://rpc.monad-testnet-2.fastlane.xyz/b3qFoDfY9sR44yRyOeHAfyj9dpEXVoOC
Bogota: https://rpc.monad-testnet-3.fastlane.xyz/j6EsEZHfw9Iqrp7DUX5e1aLUk85d1Yzw
Singapore: https://rpc.monad-testnet-5.fastlane.xyz/FFHEYATiGl2Q83xOnf71ltrqZ57q9U1W
```

## Common Commands

### Run the bot

```sh
make run-bot
# or
go run cmd/bot/main.go
```

### Check current scores

```sh
make scores
# or
go run cmd/scores/main.go
```

## Troubleshooting

If you encounter any issues:

1. Make sure Go is properly installed:

```sh
go version
```

2. Verify your `.env` file exists and contains valid credentials
3. Ensure all dependencies are installed:

```sh
go mod tidy
```

## Need Help?

- Ask for help in the FastLane on Monad Discord (#frontunner channel)
- Talk to ChatGPT
- Create an issue in this repository
- Make sure your Go environment variables are set correctly:

```sh
go env
```

## License

MIT
