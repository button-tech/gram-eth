package dto

type ExchangeTonToEthereum struct {
	ToAddress string  `json:"toAddress"`
	ToAmount  float64 `json:"toAmount"`
}

type ExchangeEthereumToTon struct {
	ToAddress string  `json:"toAddress"`
	ToAmount  float64 `json:"toAmount"`
}
