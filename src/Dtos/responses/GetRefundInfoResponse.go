package responses

type GetRefundInfoResponse struct {
	ID                         int    `json:"id"`
	RefundReference            string `json:"refundReference"`
	ParentPaymentId            int    `json:"parentPaymentId"`
	ParentTransactionReference string `json:"parentTransactionReference"`
	RefundType                 string `json:"refundType"`
	RefundAmount               int    `json:"refundAmount"`
	CreatedDate                int64  `json:"createdDate"`
	MerchantCustomerId         string `json:"merchantCustomerId"`
	RefundStatus               string `json:"refundStatus"`
	CreatedBy                  string `json:"createdBy"`
	Validated                  bool   `json:"validated"`
}
