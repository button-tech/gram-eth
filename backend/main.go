package main

import (
	"github.com/button-tech/gram-eth/backend/services/binance"
	"github.com/button-tech/gram-eth/backend/dto/ton"
	"github.com/button-tech/gram-eth/backend/services/apiClient"
	"github.com/button-tech/gram-eth/backend/services/listener"
	"log"
	"os"

	"github.com/button-tech/gram-eth/backend/services/sender"
	"github.com/button-tech/gram-eth/backend/services/server"
	"github.com/gin-gonic/gin"
)

var (
	ServerInstance *gin.Engine
)

func init() {
	endpoint := os.Getenv("ENDPOINT")
	contractAddress := os.Getenv("CONTRACT_ADDRESS")
	privateKey := os.Getenv("PRIVATE_KEY")
	tonApiUrl := os.Getenv("TON_API_URL")
	webhookUrl := os.Getenv("WEBHOOK_URL")

	txGram := ton.TxGram{
		SenderId:  os.Getenv("TON_CATALOG"),
		SenderPub: os.Getenv("TON_SENDER_PUB"),
		Network:   "0",
	}

	ServerInstance = server.InitServer()
	listener.InitListener(endpoint, contractAddress)
	sender.InitEthereum(endpoint, privateKey)
	apiClient.InitApiClient(tonApiUrl)
	go binance.ListenAndSay()
	apiClient.InitTonApiCLient(webhookUrl, &txGram)

}

func main() {

	if err := server.RunServer(ServerInstance); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
