package main

import (
	"context"
	"fmt"
	"ether/lib/blockchain"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	infuraUrl = "https://mainnet.infura.io/v3/2b4163cd4a6140d0aa3d85e66557c500"
	ganacheUrl = "http://127.0.0.1:8545"
)

func main() {
	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, ganacheUrl)
	HandleError(err)
	defer client.Close()

	block := blockchain.GetBlockByNumber(client, ctx, nil)
	fmt.Printf("%v\n", block)
}

func HandleError(err error){
	if err != nil {
		fmt.Println(err)
	}
}