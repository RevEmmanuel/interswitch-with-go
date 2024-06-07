package cardPaymentServiceRequests

type PayWithTransferRequest struct {
	MerchantCode         string `json:"merchantCode"`
	PayableCode          string `json:"payableCode"`
	CurrencyCode         string `json:"currencyCode"`
	Amount               string `json:"amount"`
	AccountName          string `json:"accountName"`
	TransactionReference string `json:"transactionReference"`
}
