package ton

type PrepareTransactionResponse struct {
	Success bool `json:"success"`
}

type CreateTransactionResponse struct {
	Success         bool   `json:"success"`
	Network         string `json:"network"`
	SenderPub       string `json:"senderPub"`
	RecipientPub    string `json:"recipientPub"`
	Amount          string `json:"amount"`
	SenderTxHash    string `json:"senderTxHash"`
	RecipientTxHash string `json:"RecipientTxHash"`
}
