package handlers

import (
	"fmt"
	"github.com/button-tech/gram-eth/backend/dto"
	"github.com/button-tech/gram-eth/backend/dto/ton"
	"github.com/button-tech/gram-eth/backend/services/apiClient"
	"github.com/button-tech/gram-eth/backend/services/sender"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CheckAuth(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

func ExchangeTonToEthereum(c *gin.Context) {
	var body dto.ExchangeTonToEthereum
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}

	value, err := strconv.ParseFloat(body.Value, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	txHash, err := sender.Send(body.SenderEthAddress, value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}

	c.JSON(http.StatusOK, dto.TxHash{
		TxHash: txHash,
	})
}

func PrepareExchangeEthereumToTon(c *gin.Context) {
	var body ton.CreateTransactionRequest
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}

	response, err := apiClient.CreateTransaction(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}

	c.JSON(http.StatusOK, response)
}

func ExchangeEthereumToTon(c *gin.Context) {
	var body ton.CreateTransactionResponse
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}

	fmt.Println(body)
	/*if !body.Success {
	apiClient.CreateTransaction(ton.CreateTransactionRequest{
		RecipientPub: body.RecipientPub,
		Amount: body.Amount,
	})
	}*/
}
