package ton

type CreateTransactionRequest struct {
	SenderId     string `json:"senderId"`
	SenderPub    string `json:"senderPub"`
	RecipientPub string `json:"recipientPub"`
	Amount       string `json:"amount"`
	Network      string `json:"network"`
	WebHookUrl   string `json:"webHookUrl"`
}

type TxGram struct {
	SenderId     string `json:"senderId"`
	SenderPub    string `json:"senderPub"`
	RecipientPub string `json:"recipientPub"`
	Amount       string `json:"amount"`
	Network      string `json:"network"`
}
