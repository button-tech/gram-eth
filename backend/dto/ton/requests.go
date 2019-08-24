package ton

type CreateTransactionRequest struct {
	SenderId     string `json:"senderId"`
	SenderPub    string `json:"SenderPub"`
	RecipientPub string `json:"recipientPub"`
	Amount       string `json:"amount"`
	Network      string `json:"network"`
	WebhookUrl   string `json:"webhookUrl"`
}
