package main

import (
	"context"
	"math/big"
	"os"
	"os/signal"
	"time"

	"github.com/FastLane-Labs/break-monad-frontrunner-bot/contract/frontrunner"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/sync/errgroup"
)

var (
	frontrunnerContractAddress = common.HexToAddress("0xBce2C725304e09CEf4cD7639760B67f8A0Af5bc4")
	taskInterval               = 1 * time.Second
)

func main() {
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, log.LvlDebug, true)))

	ethClient, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		log.Error("failed to connect to Ethereum client", "error", err)
		os.Exit(1)
	}

	chainId, err := ethClient.ChainID(context.TODO())
	if err != nil {
		log.Error("failed to get chain ID", "error", err)
		os.Exit(1)
	}

	pk, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Error("failed to parse private key", "error", err)
		os.Exit(1)
	}

	executor, err := bind.NewKeyedTransactorWithChainID(pk, chainId)
	if err != nil {
		log.Error("failed to create executor", "error", err)
		os.Exit(1)
	}

	frontrunnerContract, err := frontrunner.NewFrontrunner(frontrunnerContractAddress, ethClient)
	if err != nil {
		log.Error("failed to create Frontrunner contract", "error", err)
		os.Exit(1)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go runBot(executor, frontrunnerContract, ethClient)

	<-interrupt
}

func runBot(executor *bind.TransactOpts, frontrunnerContract *frontrunner.Frontrunner, ethClient *ethclient.Client) {
	ticker := time.NewTicker(taskInterval)
	defer ticker.Stop()

	for range ticker.C {
		feeCap, tipCap, err := getGasPrice(ethClient)
		if err != nil {
			log.Error("failed to get gas price", "error", err)
			continue
		}

		executor.GasFeeCap = feeCap
		executor.GasTipCap = tipCap

		log.Info("sending frontrun", "maxFeePerGas", feeCap, "maxPriorityFeePerGas", tipCap)

		tx, err := frontrunnerContract.Frontrun(executor)
		if err != nil {
			log.Error("failed to frontrun", "error", err)
			continue
		}

		log.Info("frontrun sent, waiting for confirmation", "hash", tx.Hash().Hex())

		receipt, err := bind.WaitMined(context.TODO(), ethClient, tx)
		if err != nil {
			log.Error("failed to land frontrun", "error", err)
			continue
		}

		log.Info("confirmed frontrun", "reverted", receipt.Status == types.ReceiptStatusFailed)
	}
}

// Returns calculated (maxFeePerGas, maxPriorityFeePerGas), override with better logic
func getGasPrice(ethClient *ethclient.Client) (*big.Int, *big.Int, error) {
	var (
		feeCap *big.Int
		tipCap *big.Int
		g, _   = errgroup.WithContext(context.TODO())
	)

	g.Go(func() error {
		_feeCap, err := ethClient.SuggestGasPrice(context.TODO())
		if err != nil {
			return err
		}

		feeCap = _feeCap
		return nil
	})

	g.Go(func() error {
		_tipCap, err := ethClient.SuggestGasTipCap(context.TODO())
		if err != nil {
			return err
		}

		tipCap = _tipCap
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, nil, err
	}

	return feeCap, tipCap, nil
}
