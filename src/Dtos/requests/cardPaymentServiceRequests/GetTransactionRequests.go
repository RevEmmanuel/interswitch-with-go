package cardPaymentServiceRequests

type GetTransactionsRequest struct {
	MerchantCode         string `json:"merchantCode" validate:"required"`
	TransactionReference string `json:"transactionReference,omitempty"`
	StartDate            string `json:"startDate" validate:"required"`
	EndDate              string `json:"endDate" validate:"required"`
	PageSize             int32  `json:"pageSize" validate:"required"`
	PageNum              int32  `json:"pageNum" validate:"required"`
}
