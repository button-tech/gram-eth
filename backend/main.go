package main

import (
	"github.com/button-tech/gram-eth/backend/services/apiClient"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/button-tech/gram-eth/backend/services/sender"
	"github.com/button-tech/gram-eth/backend/services/server"
)

var (
	ServerInstance *gin.Engine
)

func init() {
	endpoint := os.Getenv("ENDPOINT")
	privateKey := os.Getenv("PRIVATE_KEY")
	tonApiUrl := os.Getenv("TON_API_URL")
	webhookUrl := os.Getenv("WEBHOOK_URL")

	ServerInstance = server.InitServer()
	sender.InitEthereum(endpoint, privateKey)
	apiClient.InitApiClient(tonApiUrl)
	apiClient.InitTonApiCLient(webhookUrl)
}

func main() {

	if err := server.RunServer(ServerInstance); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
