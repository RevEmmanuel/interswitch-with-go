package responses

type Transaction struct {
	ID                          int64   `json:"id"`
	PayableID                   int64   `json:"payableId"`
	Amount                      int64   `json:"amount"`
	Surcharge                   int64   `json:"surcharge"`
	TransactionReference        string  `json:"transactionReference"`
	ResponseCode                string  `json:"responseCode"`
	PayableCode                 string  `json:"payableCode"`
	PayableName                 string  `json:"payableName"`
	MerchantCode                string  `json:"merchantCode"`
	MerchantCustomerName        *string `json:"merchantCustomerName"`
	Channel                     string  `json:"channel"`
	CurrencyCode                string  `json:"currencyCode"`
	PaymentCancelled            bool    `json:"paymentCancelled"`
	ChannelTransactionReference string  `json:"channelTransactionReference"`
	MerchantCustomerId          string  `json:"merchantCustomerId"`
	BankCode                    string  `json:"bankCode"`
	MaskedCardPan               string  `json:"maskedCardPan"`
	HashedCardPan               string  `json:"hashedCardPan"`
	EncryptedPan                string  `json:"encryptedPan"`
	EncryptedAccountNumber      *string `json:"encryptedAccountNumber"`
	SettlementStatus            string  `json:"settlementStatus"`
	RemittanceStatus            string  `json:"remittanceStatus"`
	DateOfPayment               string  `json:"dateOfPayment"`
	BankName                    string  `json:"bankName"`
	Status                      string  `json:"status"`
	RemittanceAmount            int64   `json:"remittanceAmount"`
	ResponseDescription         string  `json:"responseDescription"`
	CardTypeCode                string  `json:"cardTypeCode"`
	AcquirerCode                string  `json:"acquirerCode"`
	AcquirerInstitutionCode     *string `json:"acquirerInstitutionCode"`
	TransactionAuthorizationId  *string `json:"transactionAuthorizationId"`
	RetrievalReferenceNumber    *string `json:"retrievalReferenceNumber"`
	Stan                        *string `json:"stan"`
	TerminalId                  string  `json:"terminalId"`
	FeeType                     string  `json:"feeType"`
	MerchantName                string  `json:"merchantName"`
	RemittanceAccountNo         string  `json:"remittanceAccountNo"`
	RemittanceAccountType       string  `json:"remittanceAccountType"`
	RemittanceBankCode          string  `json:"remittanceBankCode"`
	NonCardProviderId           *string `json:"nonCardProviderId"`
	SplitId                     *string `json:"splitId"`
	RefundValidity              *string `json:"refundValidity"`
	ProcessedBy                 *string `json:"processedBy"`
	ContainsSplit               *bool   `json:"containsSplit"`
}

type GetTransactionsResponse struct {
	Count   int           `json:"count"`
	Content []Transaction `json:"content"`
}
