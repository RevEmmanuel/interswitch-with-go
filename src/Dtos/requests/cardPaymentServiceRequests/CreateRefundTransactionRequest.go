package cardPaymentServiceRequests

type CreateRefundTransactionRequest struct {
	RefundReference            string `json:"refundReference"`
	ParentTransactionReference string `json:"parentTransactionReference"`
	ParentPaymentId            string `json:"parentPaymentId"`
	RefundType                 string `json:"refundType"`
	RefundAmount               string `json:"refundAmount"`
	CreatedBy                  string `json:"createdBy"`
	ReasonsForRefund           string `json:"reasonsForRefund"`
}
