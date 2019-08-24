package listener

import (
	"context"
	"fmt"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/button-tech/gram-eth/backend/services/ethereum"
	"log"
	"math/big"
	"strings"
	"time"
)

var (
	Contract *ethereum.Contract
	Eth      *ethereum.SingleTransaction

	transactionHashes []string
)

func InitEthereum(endpoint, privateKey string) {
	var err error
	Eth, err = ethereum.Connect(endpoint, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	go listen()
}

func listen() {
	query := eth.FilterQuery{
		Addresses: []common.Address{Contract.Address},
	}

	contractAbi, err := abi.JSON(strings.NewReader(string("abi")))
	if err != nil {
		log.Fatal(err)
	}

	logs := make(chan types.Log)
	sub, err := Contract.Client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			event := struct {
				TonAddress string
				NewRate   *big.Int
			}{}
			err := contractAbi.Unpack(&event, "Received", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("TON Address: " + event.TonAddress)
			fmt.Println("TON Value: " + event.NewRate.String())

		}
		time.Sleep(3 * time.Second)
	}
}
