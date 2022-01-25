package blockchain

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)



func GetBlockByNumber(client *ethclient.Client, ctx context.Context, blockNumber *big.Int) *types.Block{
	block, err := client.BlockByNumber(ctx, blockNumber)
	HandleError(err)

	return block
}


func HandleError(err error){
	if err != nil {
		fmt.Println(err)
	}
}