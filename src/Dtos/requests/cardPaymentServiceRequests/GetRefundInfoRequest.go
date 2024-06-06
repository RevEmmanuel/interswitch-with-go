package cardPaymentServiceRequests

type GetRefundInfoRequest struct {
	RefundReference string `json:"refundReference" binding:"required"`
}
