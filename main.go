package main

import (
	"context"
	"ether/lib/blockchain"
	"ether/lib/wallet"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	infuraUrl = "https://mainnet.infura.io/v3/2b4163cd4a6140d0aa3d85e66557c500"
	ganacheUrl = "http://127.0.0.1:8545"
	walletAddress = "0x6eA71CE756BCAa7536e765B0Aa15Cb1aE706EF7e"
)

func main() {
	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, infuraUrl)
	HandleError(err)
	defer client.Close()

	_ = blockchain.GetBlockByNumber(client, ctx, nil)

	address := common.HexToAddress(walletAddress)

	res := wallet.GetBalance(client, ctx, address, nil)

	balance := new(big.Float)
	balance.SetString(res.String())



	value := new(big.Float).Quo(balance, big.NewFloat(math.Pow10(18)))

	fmt.Println(value)
}

func HandleError(err error){
	if err != nil {
		fmt.Println(err)
	}
}