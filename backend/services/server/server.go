package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	. "github.com/button-tech/gram-eth/backend/services/server/handlers"
)

func InitServer() *gin.Engine {
	g := gin.New()
	gin.SetMode(gin.ReleaseMode)
	return g
}

func RunServer(R *gin.Engine) error {
	R.Use(gin.Recovery())
	R.Use(gin.Logger())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")
	R.Use(cors.New(config))
	{

		api := R.Group("/api")
		{
			api.POST("/ton/eth", ExchangeTonToEthereum)
			api.POST("/eth/ton/prepare", PrepareExchangeEthereumToTon)
			api.POST("/eth/ton", ExchangeEthereumToTon)
		}
	}

	if err := R.Run(":8080"); err != nil {
		return err
	}

	return nil
}
