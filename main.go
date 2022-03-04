package main

import (
	"fmt"
	"log"
	"math/big"

	"box/api"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	contractAddress := common.HexToAddress("0x4bC93B24f371dB2E1Cae9877d147634807689CB1")
	conn, err := api.NewApi(contractAddress, client)
	if err != nil {
		panic(err)
	}
	// query := ethereum.FilterQuery{
	// 	Addresses: []common.Address{contractAddress},
	// }

	// logs := make(chan types.Log)
	// sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	// if err != nil {
	// 	log.Println("Faile to sucribe fileter logs", err)
	// }

	// fmt.Println("printing subsribe", sub)
	reply, err := conn.Store(&bind.TransactOpts{}, big.NewInt(100))
	if err != nil {
		fmt.Println("Faile to store value", err.Error())
	}
	fmt.Println("Printing reply", reply)

	// reply2, err := conn.Retrieve(&bind.CallOpts{})
	// if err != nil {
	// 	fmt.Println("failed to retrieve value", err.Error())
	// }
	// fmt.Println("Printing reply2", reply2)

	// for {
	// 	select {
	// 	case err := <-sub.Err():
	// 		log.Fatal(err)
	// 	case vLog := <-logs:
	// 		fmt.Println(vLog) // pointer to event log
	// 	}
	// }

}
