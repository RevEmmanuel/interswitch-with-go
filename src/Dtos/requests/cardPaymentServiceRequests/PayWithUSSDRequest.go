package cardPaymentServiceRequests

type PayWithUSSDRequest struct {
	Amount                       string `json:"amount"`
	BankCode                     string `json:"bankCode"`
	Surcharge                    string `json:"surcharge"`
	CurrencyCode                 string `json:"currencyCode"`
	MerchantTransactionReference string `json:"merchantTransactionReference"`
}
