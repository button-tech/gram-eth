package listener

import (
	"context"
	"fmt"
	"github.com/button-tech/gram-eth/backend/contract"
	"github.com/button-tech/gram-eth/backend/dto/ton"
	"github.com/button-tech/gram-eth/backend/services/apiClient"
	"github.com/button-tech/gram-eth/backend/services/ethereum"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"
)

var (
	Contract *ethereum.Contract

	transactionHashes []string
)

func InitListener(endpoint, contractAddress string) {
	var err error
	Contract, err = ethereum.ConnectContract(endpoint, contractAddress)
	if err != nil {
		log.Fatal(err)
	}
	go listen()
}

func listen() {
	query := eth.FilterQuery{
		Addresses: []common.Address{Contract.Address},
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(contract.ContractABI)))
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
				SumToSend  *big.Int
			}{}
			err := contractAbi.Unpack(&event, "EtherRecieved", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("TON Address: " + event.TonAddress)
			fmt.Println("TON Value: " + event.SumToSend.String())

			amountLog := event.SumToSend.Int64()

			bigAmountLog := new(big.Float).SetFloat64(float64(amountLog))

			multiplier := new(big.Float).SetFloat64(math.Pow(10, 16))

			var valueToSend string
			sendingAmount, _ := new(big.Float).Quo(bigAmountLog, multiplier).Float64()

			str := strconv.FormatFloat(sendingAmount, 'g', 1, 64)
			if float64(int(sendingAmount))-sendingAmount == 0 {
				valueToSend = strings.Split(str, "e+")[0] + ".0"
			} else {
				valueToSend = strings.Split(str, "e+")[0]
			}

			response, err := apiClient.CreateTransaction(ton.CreateTransactionRequest{
				Amount:       valueToSend,
				RecipientPub: event.TonAddress,
			})
			if err != nil {
				log.Println(err)
				return
			}
			if response.Success {
				log.Println("true")
			}

		}
		time.Sleep(3 * time.Second)
	}
}
