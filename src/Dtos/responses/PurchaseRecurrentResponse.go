package responses

type PurchaseRecurrentResponse struct {
	TransactionRef          string `json:"transactionRef"`
	PaymentId               string `json:"paymentId"`
	BankCode                string `json:"bankCode"`
	Message                 string `json:"message"`
	Amount                  string `json:"amount"`
	ResponseCode            string `json:"responseCode"`
	PlainTextSupportMessage string `json:"plainTextSupportMessage"`
}
