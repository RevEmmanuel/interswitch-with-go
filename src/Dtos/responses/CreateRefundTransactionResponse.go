package responses

type CreateRefundTransactionResponse struct {
	Id                         int    `json:"id"`
	RefundReference            string `json:"refundReference"`
	MerchantCode               string `json:"merchantCode"`
	ParentPaymentId            int    `json:"parentPaymentId"`
	ParentTransactionReference string `json:"parentTransactionReference"`
	RefundType                 string `json:"refundType"`
	RefundAmount               int    `json:"refundAmount"`
	CreatedDate                int64  `json:"createdDate"`
	RefundStatus               string `json:"refundStatus"`
	CreatedBy                  string `json:"createdBy"`
	Validated                  bool   `json:"validated"`
}
