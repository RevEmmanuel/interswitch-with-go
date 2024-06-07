package cardPaymentServiceRequests

type GenerateQRRequest struct {
	Amount                       string `json:"amount"`
	Surcharge                    string `json:"surcharge"`
	CurrencyCode                 string `json:"currencyCode"`
	MerchantTransactionReference string `json:"merchantTransactionReference"`
}
