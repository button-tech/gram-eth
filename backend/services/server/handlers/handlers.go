package handlers

import (
	"fmt"
	"github.com/button-tech/gram-eth/backend/dto"
	"github.com/button-tech/gram-eth/backend/dto/ton"
	"github.com/button-tech/gram-eth/backend/services/apiClient"
	"github.com/button-tech/gram-eth/backend/services/sender"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strconv"
)

func CheckAuth(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

func ExchangeTonToEthereum(c *gin.Context) {
	var body dto.ExchangeTonToEthereum
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}

	floatAmount, err := strconv.ParseFloat(body.Value, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}
	bigFloatAmount := new(big.Float).SetFloat64(floatAmount)
	value, _ := new(big.Float).Mul(bigFloatAmount, new(big.Float).SetFloat64(0.001)).Float64()


	amount := strconv.FormatFloat(value, 'g', 1, 64)
	apiClient.Nitify(dto.TonNotify{
		TonPubAddress: body.From,
		Description: "💎 Your " + amount + " ETH are on the way!",
	})

	txHash, err := sender.Send(body.SenderEthAddress, value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}

	c.JSON(http.StatusOK, dto.TxHash{
		TxHash: txHash,
	})
}

func ExchangeTonToBnb(c *gin.Context) {
/*	var body dto.ExchangeTonToEthereum
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}

	floatAmount, err := strconv.ParseFloat(body.Value, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}
	bigFloatAmount := new(big.Float).SetFloat64(floatAmount)
	value, _ := new(big.Float).Mul(bigFloatAmount, new(big.Float).SetFloat64(10^6)).Int64()*/

	//binance.SendTransaction(body.SenderEthAddress, value)
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
	err = apiClient.Nitify(dto.TonNotify{
		TonPubAddress: body.RecipientPub,
		Description: "💎 Your " + body.Amount + " GRAMs are on the way! We will notify you when it will be done.",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}
	c.JSON(http.StatusOK, response)
}

func PrepareExchangeWavesToTon(c *gin.Context) {
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
	err = apiClient.Nitify(dto.TonNotify{
		TonPubAddress: body.RecipientPub,
		Description: "💎 Your " + body.Amount + " GRAMs are on the way! We will notify you when it will be done.",
	})
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
	if !body.Success {
		apiClient.Nitify(dto.TonNotify{
			TonPubAddress: body.RecipientPub,
			Description: "❌ Exchange failed.",
		})
		return
	}

	apiClient.Nitify(dto.TonNotify{
		TonPubAddress: body.RecipientPub,
		Description: "💎 Successfully got " + body.Amount + " GRAMs!",
	})
}
