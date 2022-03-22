package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"box/api"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// client, err := ethclient.Dial("http://127.0.0.1:8545")
	client, err := ethclient.Dial("ws://localhost:7545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	contractAddress := common.HexToAddress("0xbf8d8e47B960cB5eDC3e4274eccC5fB6DEDb65b4")
	conn, err := api.NewApi(contractAddress, client)
	if err != nil {
		fmt.Println(err)
	}

	privateKey, err := crypto.HexToECDSA("37b7bd4ef5c92a3f74e708a2cc070997ce5784fa21adeb9c971d9907daea1e78")
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000)
	opts := &bind.CallOpts{}
	gold, err := conn.SILVER(opts)

	fmt.Println("Printing gold", gold)
	// ch := make(chan *api.ApiValueChanged)

	// // opts.Start = &blockNumber
	// sub, err := conn.WatchValueChanged(opts, ch)
	// // _, err := contract.W(opts, ch, s)
	// if err != nil {
	// 	log.Fatalf("Failed WatchYearChanged: %v", err)
	// }

	// go events(sub, ch)
	// _, err = conn.Store(&bind.TransactOpts{
	// 	From:   fromAddress,
	// 	Signer: auth.Signer,
	// }, big.NewInt(1100))
	// if err != nil {
	// 	fmt.Println("Faile to store value", err.Error())
	// }

	// _, err = conn.Retrieve(&bind.CallOpts{})
	// if err != nil {
	// 	fmt.Println("failed to retrieve value", err.Error())
	// }

}

// func events(sub event.Subscription, ch <-chan *api.ApiValueChanged) {

// 	for {
// 		select {
// 		case err := <-sub.Err():
// 			log.Fatal(err)
// 		case vLog := <-ch:
// 			fmt.Println(vLog.Value) // pointer to event log
// 			fmt.Println("Printing vlog", vLog)
// 			// fmt.Println("I was here", vLog.Data)
// 		}
// 	}
// }
