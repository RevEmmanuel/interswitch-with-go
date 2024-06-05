package responses

type TokenizeCardResponse struct {
	TransactionRef  string `json:"transactionRef"`
	AccountNumber   string `json:"accountNumber"`
	Token           string `json:"token"`
	TokenExpiryDate string `json:"tokenExpiryDate"`
	PanLast4Digits  string `json:"panLast4Digits"`
	CardType        string `json:"cardType"`
	Balance         string `json:"balance"`
}
