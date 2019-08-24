package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/artall64/gram-eth/backend/dto"
	"github.com/artall64/gram-eth/backend/services/sender"
	"net/http"
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

	txHash, err := sender.Send(body.ToAddress, body.ToAmount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}

	c.JSON(http.StatusOK, dto.TxHash{
		TxHash: txHash,
	})
}
