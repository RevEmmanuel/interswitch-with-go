package cardPaymentServiceRequests

type GetRefundRequest struct {
	MerchantCode string `json:"merchantCode" binding:"required"`
	PageNum      string `json:"pageNum,omitempty"`
	PageSize     string `json:"pageSize,omitempty"`
	StartDate    string `json:"startDate,omitempty"`
	EndDate      string `json:"endDate,omitempty"`
}
