package main

import (
	"context"
	// "crypto"
	"ether/lib/blockchain"
	"ether/lib/wallet"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ganacheUrl = "http://127.0.0.1:8545"
	address = common.HexToAddress("0xD0F4752e86c9d207a400caFbb973b80E82CCE559")
	ctx = context.Background()
	client, _ = ethclient.DialContext(ctx, ganacheUrl)
	privKey, _ = crypto.GenerateKey()
	privateKey = hexutil.Encode(crypto.FromECDSA(privKey))
	publicKey = hexutil.Encode(crypto.FromECDSAPub(&privKey.PublicKey))
	Address = crypto.PubkeyToAddress(privKey.PublicKey).Hex()
)

func main() {
	defer client.Close()
	// GetBlock()
	// GetBalance()
	fmt.Println(publicKey)
	fmt.Println(privateKey)
	fmt.Println(Address)
}

func HandleError(err error){
	if err != nil {
		fmt.Println(err)
	}
}


func GetBlock() {
	block := blockchain.GetBlockByNumber(client, ctx, nil)
	fmt.Println(block.Number())
}

func GetBalance() {
	res := wallet.GetBalance(client, ctx, address, nil)

	balance := new(big.Float)
	balance.SetString(res.String())



	value := new(big.Float).Quo(balance, big.NewFloat(math.Pow10(18)))

	fmt.Println(value)
}