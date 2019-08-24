package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/artall64/gram-eth/backend/services/sender"
	"github.com/artall64/gram-eth/backend/services/server"
)

var (
	ServerInstance *gin.Engine
)

func init() {
	ServerInstance = server.InitServer()
	endpoint := os.Getenv("ENDPOINT")
	privateKey := os.Getenv("PRIVATE_KEY")
	sender.InitEthereum(endpoint, privateKey)
}

func main() {

	if err := server.RunServer(ServerInstance); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
