package notification

import (
	"fmt"
	"github.com/imroc/req"
)

type R struct{
	PubAddr string `json:"tonPubAddress"`
	Text string `json:"description"`
}

func SendNotification(tonAddress string) {
	url := "https://client.buttonwallet.com/api/tontestnet/notify"

	header := req.Header{
		"Content-Type": "application/json",
	}
	var payload  = R{PubAddr:tonAddress, Text:"📲 New TON transaction will come soon to your account"}
	response, err := req.Post(url, header, payload)
	if err != nil {
		fmt.Println(err)
	}

	code := response.Response().StatusCode
	fmt.Println(code)
}
