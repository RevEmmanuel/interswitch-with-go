package cardPaymentServiceRequests

type PurchaseRecurrentRequest struct {
	CustomerId      string `json:"customerId"`
	Amount          string `json:"amount"`
	Currency        string `json:"currency"`
	Token           string `json:"token"`
	TokenExpiryDate string `json:"tokenExpiryDate"`
	TransactionRef  string `json:"transactionRef"`
}
