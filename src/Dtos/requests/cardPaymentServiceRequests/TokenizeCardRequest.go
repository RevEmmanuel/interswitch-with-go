package cardPaymentServiceRequests

type TokenizeCardRequest struct {
	TransactionRef string `json:"transactionRef"`
	AuthData       string `json:"authData"`
}
