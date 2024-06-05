package responses

import "time"

type TokenizeCardResponse struct {
	CardType        string    `json:"cardType"`
	Balance         string    `json:"balance"`
	Token           string    `json:"token"`
	TokenExpiryDate time.Time `json:"tokenExpiryDate"`
	PanLast4Digits  string    `json:"panLast4Digits"`
	TransactionRef  string    `json:"transactionRef"`
}
