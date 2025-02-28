package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"
	binding "temp-contracts/src"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"
)

var (
	app *cli.App
)

func init() {
	app = cli.NewApp()
	app.Commands = []*cli.Command{
		{
			Name:   "deploy",
			Usage:  "deploy PrevRandao Contract",
			Action: deploy,
			Flags: []cli.Flag{
				&rpcURL,
				&privateKey,
			},
		},
		{
			Name:   "reset",
			Usage:  "Run reset",
			Action: run,
			Flags: []cli.Flag{
				&rpcURL,
				&privateKey,
				&address,
			},
		},
	}
	app.Flags = []cli.Flag{}
}

var (
	rpcURL = cli.StringFlag{
		Name:     "rpc-url",
		Required: true,
	}
	privateKey = cli.StringFlag{
		Name:     "private-key",
		Required: true,
	}
	address = cli.StringFlag{
		Name:     "address",
		Required: true,
	}
)

func main() {
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func deploy(ctx *cli.Context) error {
	rpcurl := ctx.String(rpcURL.Name)
	client, err := ethclient.DialContext(ctx.Context, rpcurl)
	if err != nil {
		return err
	}
	chainID, err := client.ChainID(ctx.Context)
	if err != nil {
		return err
	}
	fmt.Println("chainID: ", chainID)

	pk, err := crypto.HexToECDSA(strings.TrimPrefix(ctx.String(privateKey.Name), "0x"))
	if err != nil {
		return err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		return err
	}
	auth.GasLimit = 1e7

	address, tx, _, err := binding.DeployPrevRandao(auth, client)
	if err != nil {
		return err
	}
	receipt, err := bind.WaitMined(ctx.Context, client, tx)
	if err != nil {
		return err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("transaction failed")
	}
	fmt.Println("PrevRandao Address: ", address.Hex())

	return nil
}

func run(ctx *cli.Context) error {
	rpcurl := ctx.String(rpcURL.Name)
	client, err := ethclient.DialContext(ctx.Context, rpcurl)
	if err != nil {
		return err
	}
	chainID, err := client.ChainID(ctx.Context)
	if err != nil {
		return err
	}
	fmt.Println("chainID: ", chainID)

	pk, err := crypto.HexToECDSA(strings.TrimPrefix(ctx.String(privateKey.Name), "0x"))
	if err != nil {
		return err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		return err
	}
	auth.GasLimit = 1e7
	nonce, err := client.NonceAt(context.Background(), auth.From, nil)
	if err != nil {
		return err
	}
	auth.Nonce = new(big.Int).SetUint64(nonce)

	prevrandao := common.HexToAddress(ctx.String(address.Name))
	PrevRandao, err := binding.NewPrevRandao(prevrandao, client)
	if err != nil {
		return err
	}

	callOpts := &bind.CallOpts{Context: ctx.Context}
	for {
		tx, err := PrevRandao.Reset(auth)
		if err != nil {
			fmt.Println("send transaction err", err)
		} else {
			auth.Nonce.Add(auth.Nonce, common.Big1)
		}
		if tx != nil {
			go func(tx *types.Transaction) {
				receipt, err := bind.WaitMined(ctx.Context, client, tx)
				if err != nil {
					fmt.Println("wait receipt err", err)
				}
				if receipt.Status != types.ReceiptStatusSuccessful {
					fmt.Println("tx reverted", receipt.TxHash)
					return
				}
				value, err := PrevRandao.Value(callOpts)
				if err != nil {
					fmt.Println("call value err", err)
				}
				number, err := PrevRandao.Blocknumber(callOpts)
				if err != nil {
					fmt.Println("call blockNumber", err)
				}
				fmt.Printf("receipt.BlockNumber: %d\tcall.BlockNumber: %d\tvalue: %d\n", receipt.BlockNumber, number, value)
			}(tx)
		}
		time.Sleep(1e9)
	}
}
