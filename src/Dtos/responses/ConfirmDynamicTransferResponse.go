package responses

type ConfirmDynamicTransferResponse struct {
	ResponseCode                string  `json:"responseCode"`
	ResponseDescription         string  `json:"responseDescription"`
	TransactionReference        string  `json:"transactionReference"`
	ChannelTransactionReference string  `json:"channelTransactionReference"`
	Amount                      int     `json:"amount"`
	RemittanceAmount            int     `json:"remittanceAmount"`
	CustomerName                *string `json:"customerName,omitempty"`
	Bank                        string  `json:"bank"`
	AccountNumber               string  `json:"accountNumber"`
}
