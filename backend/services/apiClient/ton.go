package apiClient

import (
	"github.com/imroc/req"
	"github.com/button-tech/gram-eth/backend/dto/ton"
	"net/http"
)

var (
	webhookUrl string
)

func InitTonApiCLient(webhookUrl string) {
	webhookUrl = webhookUrl
}

func CreateTransaction(request ton.CreateTransactionRequest) (ton.PrepareTransactionResponse, error) {
	request.WebhookUrl = webhookUrl
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
