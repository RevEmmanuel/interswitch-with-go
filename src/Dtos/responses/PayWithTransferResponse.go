package responses

type PayWithTransferResponse struct {
	AccountNumber        string `json:"accountNumber"`
	BankName             string `json:"bankName"`
	Amount               int    `json:"amount"`
	TransactionReference string `json:"transactionReference"`
	ResponseCode         string `json:"responseCode"`
	ValidityPeriodMins   int    `json:"validityPeriodMins"`
	AccountName          string `json:"accountName"`
}
