package sender

import (
	"github.com/artall64/gram-eth/backend/services/ethereum"
	"log"
)

var (
	Eth *ethereum.SingleTransaction
)

func InitEthereum(endpoint, privateKey string) {
	var err error
	Eth, err = ethereum.Connect(endpoint, privateKey)
	if err != nil {
		log.Fatal(err)
	}
}

func Send(address string, amount float64) (string, error) {
	toAddress := address[2:]
	toAmount := ethereum.EtherToInt(amount)

	tx, err := Eth.SignTransaction(toAddress, toAmount)
	if err != nil {
		return "", err
	}

	txHash, err := Eth.SendSignedTransaction(tx)
	if err != nil {
		return "", err
	}

	return txHash, nil
}
