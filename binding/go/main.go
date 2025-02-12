package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"
	binding "temp-contracts/src"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

func main() {
	err := tstore()
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
	}
}

func env() (*bind.TransactOpts, *ethclient.Client, error) {
	privateKey := os.Getenv("PRIVATE_KEY")
	rpc_url := os.Getenv("RPC_URL")

	client, err := ethclient.Dial(rpc_url)
	if err != nil {
		return nil, nil, err
	}

	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, nil, err
	}
	chainId, err := client.ChainID(context.TODO())
	if err != nil {
		return nil, nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainId)
	return auth, client, err
}

func tstore() error {
	auth, backend, err := env()
	if err != nil {
		return err
	}
	op, tx1, OP, err := binding.DeployCancunOp(auth, backend)
	if err != nil {
		return err
	}
	caller, tx2, CALLER, err := binding.DeployCancunOpTCaller(auth, backend, op)
	if err != nil {
		return err
	}

	ctx := context.Background()
	if receipt, err := bind.WaitMined(ctx, backend, tx1); err != nil {
		return err
	} else if receipt.Status != types.ReceiptStatusSuccessful {
		return errors.New("revert")
	}
	if receipt, err := bind.WaitMined(ctx, backend, tx2); err != nil {
		return err
	} else if receipt.Status != types.ReceiptStatusSuccessful {
		return errors.New("revert")
	}

	fmt.Println("\n", strings.Repeat("-", 80))
	bytes, err := backend.CodeAt(ctx, op, nil)
	if err != nil {
		return err
	}
	fmt.Println("OP", op.Hex(), "size", len(bytes))
	bytes, err = backend.CodeAt(ctx, caller, nil)
	if err != nil {
		return err
	}
	fmt.Println("CALLER", caller.Hex(), "size", len(bytes))

	fmt.Println("\n", strings.Repeat("-", 80))
	callOpts := &bind.CallOpts{Context: ctx}
	value, err := OP.Value(callOpts)
	if err != nil {
		return err
	}
	fmt.Println("before", value)

	setValue := big.NewInt(32)
	tx, err := OP.SetValue(auth, setValue)
	if err != nil {
		return err
	}

	if receipt, err := bind.WaitMined(ctx, backend, tx); err != nil {
		return errors.Wrap(err, "SetValue")
	} else if receipt.Status != types.ReceiptStatusSuccessful {
		return errors.New("revert SetValue")
	} else {
		event, err := OP.ParseSetValue(*receipt.Logs[0])
		if err != nil {
			return err
		} else {
			if setValue.Cmp(event.Value) != 0 {
				fmt.Println("emit SetValue", event.Value, "value", setValue)
			}
		}
	}

	value, err = OP.Value(callOpts)
	if err != nil {
		return err
	}
	fmt.Println("after", value)
	fmt.Println("\n", strings.Repeat("-", 80))
	value, err = CALLER.Value(callOpts)
	if err != nil {
		return err
	}
	fmt.Println("before", value)
	setValue = big.NewInt(64)
	tx, err = OP.Execute(auth, caller, setValue)
	if err != nil {
		return err
	}
	if receipt, err := bind.WaitMined(ctx, backend, tx); err != nil {
		return errors.Wrap(err, "Execute")
	} else if receipt.Status != types.ReceiptStatusSuccessful {
		return errors.New("revert Execute")
	} else {
		event, err := OP.ParseSetValue(*receipt.Logs[0])
		if err != nil {
			return err
		} else {
			if setValue.Cmp(event.Value) != 0 {
				fmt.Println("emit SetValue", event.Value, "value", setValue)
			}
		}
	}
	value, err = CALLER.Value(callOpts)
	if err != nil {
		return err
	}
	fmt.Println("after", value)

	return nil
}
