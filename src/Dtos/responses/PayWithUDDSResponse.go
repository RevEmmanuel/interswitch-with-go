package responses

type PayWithUSSDResponse struct {
	Reference            string `json:"reference"`
	BankShortCode        string `json:"bankShortCode"`
	TransactionReference string `json:"transactionReference"`
	ResponseCode         string `json:"responseCode"`
	DefaultShortCode     string `json:"defaultShortCode"`
}
