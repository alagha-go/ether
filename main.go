package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	infuraUrl = "https://mainnet.infura.io/v3/2b4163cd4a6140d0aa3d85e66557c500"
)

func main() {
	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, infuraUrl)
	HandleError(err)
	defer client.Close()

	block, err := client.BlockByNumber(ctx, nil)
	HandleError(err)
	fmt.Println(block.Number())
}

func HandleError(err error){
	if err != nil {
		fmt.Println(err)
	}
}