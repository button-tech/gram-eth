package apiClient

import (
	"github.com/button-tech/gram-eth/backend/dto"
	"github.com/button-tech/gram-eth/backend/dto/ton"
	"github.com/imroc/req"
	"net/http"
)

var (
	webHookUrl string
	txData     *ton.TxGram
)

func InitTonApiCLient(url string, gram *ton.TxGram) {
	webHookUrl = url
	txData = gram
}

func CreateTransaction(request ton.CreateTransactionRequest) (ton.PrepareTransactionResponse, error) {
	request.WebHookUrl = webHookUrl
	request.Network = txData.Network
	request.SenderId = txData.SenderId
	request.SenderPub = txData.SenderPub

	call := apiCall("POST", "/send", request)
	if call.StatusCode != http.StatusOK {
		return ton.PrepareTransactionResponse{}, call.Error
	}
	var responseToClient ton.PrepareTransactionResponse
	if err := call.Result.(*req.Resp).ToJSON(&responseToClient); err != nil {
		return ton.PrepareTransactionResponse{}, call.Error
	}
	return responseToClient, nil
}

func Nitify(data dto.TonNotify) (error) {
	var (
		respErr    error
		authHeader = req.Header{
			"Content-Type": "application/json",
		}
	)
	uri := "https://client.buttonwallet.com/api/tontestnet/notify"
	_, respErr = req.Post(uri, authHeader, data)
	if respErr != nil {
		return respErr
	}
	return nil
}