package cardPaymentServiceRequests

type ConfirmDynamicTransferRequest struct {
	MerchantCode         string `json:"merchantCode" binding:"required"`
	TransactionReference string `json:"transactionReference" binding:"required"`
}
