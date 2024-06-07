package responses

type Issuer struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	DomainTypeID int    `json:"domainTypeId"`
	CountryCode  string `json:"countryCode"`
	CbnCode      string `json:"cbnCode"`
}

type CardTypeGatewayConfiguration struct {
	ID                                        int    `json:"id"`
	SupportsPin                               bool   `json:"supportsPin"`
	SupportsCvv2                              bool   `json:"supportsCvv2"`
	SupportsExpiryDate                        bool   `json:"supportsExpiryDate"`
	Enabled                                   bool   `json:"enabled"`
	CardTypeCode                              string `json:"cardTypeCode"`
	SupportsOtp                               bool   `json:"supportsOtp"`
	SupportsCardinalAuthentication            bool   `json:"supportsCardinalAuthentication"`
	LoyaltyRedeemEnabled                      bool   `json:"loyaltyRedeemEnabled"`
	LoyaltyEnabled                            bool   `json:"loyaltyEnabled"`
	SupportsFingerPrintAuthorization          bool   `json:"supportsFingerPrintAuthorization"`
	SupportsAccountType                       bool   `json:"supportsAccountType"`
	SupportsVisaAcceleratedConnectionPlatform bool   `json:"supportsVisaAcceleratedConnectionPlatform"`
	SupportsCybersource20                     bool   `json:"supportsCybersource20"`
	CybersourceEnabled                        bool   `json:"cybersourceEnabled"`
	SupportThreeDsAuthentication              bool   `json:"supportThreeDsAuthentication"`
	SupportsWibmoAuthentication               bool   `json:"supportsWibmoAuthentication"`
}

type PaymentMethod struct {
	CardTypeCode                 string                       `json:"cardTypeCode"`
	CardTypeName                 string                       `json:"cardTypeName"`
	MaskedPan                    string                       `json:"maskedPan"`
	AccountNumber                string                       `json:"accountNumber"`
	CardIdentifier               string                       `json:"cardIdentifier"`
	CardHash                     string                       `json:"cardHash"`
	WalletInstrumentIdentifier   string                       `json:"walletInstrumentIdentifier"`
	Status                       string                       `json:"status"`
	ExpiryMonth                  string                       `json:"expiryMonth"`
	ExpiryYear                   string                       `json:"expiryYear"`
	Issuer                       Issuer                       `json:"issuer"`
	Name                         string                       `json:"name"`
	CookieKey                    string                       `json:"cookieKey"`
	LoyaltyRedemptionAllowed     bool                         `json:"loyaltyRedemptionAllowed"`
	MaskToken                    bool                         `json:"maskToken"`
	EnableFingerPrint            bool                         `json:"enableFingerPrint"`
	CardTypeGatewayConfiguration CardTypeGatewayConfiguration `json:"cardTypeGatewayConfiguration"`
}

type User struct {
	ID         interface{} `json:"id"`
	Username   interface{} `json:"username"`
	Email      string      `json:"email"`
	MobileNo   string      `json:"mobileNo"`
	FirstName  string      `json:"firstName"`
	LastName   string      `json:"lastName"`
	PassportID interface{} `json:"passportId"`
	Active     bool        `json:"active"`
	Admin      bool        `json:"admin"`
	DomainID   interface{} `json:"domainId"`
	DomainName interface{} `json:"domainName"`
	DomainCode interface{} `json:"domainCode"`
	RoleID     interface{} `json:"roleId"`
	Roles      interface{} `json:"roles"`
	Domains    interface{} `json:"domains"`
	Apps       interface{} `json:"apps"`
}

type GetWalletCardsResponse struct {
	PaymentMethods []PaymentMethod `json:"paymentMethods"`
	User           User            `json:"user"`
}
