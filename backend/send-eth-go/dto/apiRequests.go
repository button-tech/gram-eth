package dto

type ExchangeTonToEthereum struct {
	To               string `json:"to"`
	From             string `json:"from"`
	SenderEthAddress string `json:"senderEthAddress"`
	Guid             string `json:"guid"`
	Value            string `json:"value"`
	TxHashSender     string `json:"txHashSender"`
	TxHashRecipient  string `json:"txHashRecipient"`
}

type ExchangeEthereumToTon struct {
	ToAddress string  `json:"toAddress"`
	ToAmount  float64 `json:"toAmount"`
}
