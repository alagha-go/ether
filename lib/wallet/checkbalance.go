package wallet

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBalance(client *ethclient.Client, ctx context.Context, address common.Address, block *big.Int) *big.Int{
	balance, err := client.BalanceAt(ctx, address, nil)
	HandleError(err)

	return balance
}

func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}