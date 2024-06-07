package responses

type GenerateQRResponse struct {
	QRCodeId             string `json:"qrCodeId"`
	QRCodeIdMasterPass   string `json:"qrCodeIdMasterPass"`
	RawQRData            string `json:"rawQRData"`
	TransactionReference string `json:"transactionReference"`
}
