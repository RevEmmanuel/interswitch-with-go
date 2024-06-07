package responses

type Merchant struct {
	Name                    string `json:"name"`
	MerchantCode            string `json:"merchantCode"`
	LogoUrl                 string `json:"logoUrl"`
	CorporateApprovalStatus int    `json:"corporateApprovalStatus"`
	SettlementType          int    `json:"settlementType"`
}

type Provider struct {
	ID                    int               `json:"id"`
	PaymentLabel          string            `json:"paymentLabel"`
	RequiredPaymentFields map[string]string `json:"requiredPaymentFields"`
	SupportsCallback      bool              `json:"supportsCallback"`
	ProviderId            string            `json:"providerId"`
	CategoryCode          string            `json:"categoryCode"`
	EsbRoute              string            `json:"esbRoute"`
	Enabled               bool              `json:"enabled"`
	ProviderDetails       ProviderDetails   `json:"provider"`
	Country               string            `json:"country"`
}

type ProviderDetails struct {
	ID              int    `json:"id"`
	InstitutionCode string `json:"institutionCode"`
	ProviderId      string `json:"providerId"`
	CategoryId      int    `json:"categoryId"`
}

type PaymentOption struct {
	PayableCode           string                 `json:"payableCode"`
	ProviderCode          string                 `json:"providerCode"`
	Enabled               bool                   `json:"enabled"`
	AdditionalInformation map[string]interface{} `json:"additionalInformation"`
}

type GenerateAlternativePaymentOptionsResponse struct {
	Merchant       Merchant        `json:"merchant"`
	PaymentOptions []PaymentOption `json:"paymentOptions"`
}
