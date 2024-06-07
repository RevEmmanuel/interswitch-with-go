package cardPaymentAPITests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"interswitch_go_testing/src/Dtos/requests/cardPaymentServiceRequests"
	"interswitch_go_testing/src/credentialConfig"
	"interswitch_go_testing/src/services/cardPaymentService"
	"interswitch_go_testing/src/utils"
	"math/rand"
	"testing"
	"time"
)

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	character := make([]byte, length)
	for i := range character {
		character[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(character)
}

func TestTokenizeCardRecurrentSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	request := cardPaymentServiceRequests.TokenizeCardRequest{
		TransactionRef: credentialConfig.TRANSACTIONREF,
		AuthData:       credentialConfig.AUTHDATA,
	}

	response, err := cardPaymentService.TokenizeCardRecurrent(request)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if response == nil {
		t.Fatalf("expected response, got nil")
	} else {
		assert.Nil(t, err, "no errors")
		assert.NotEmpty(t, response.CardType)
		assert.NotEmpty(t, response.Balance)
		assert.Equal(t, response.CardType, "Verve")
	}
}

func TestTokenizeCardRecurrentFailure(t *testing.T) {
	gin.SetMode(gin.TestMode)
	request := cardPaymentServiceRequests.TokenizeCardRequest{
		TransactionRef: "invalid_transactionRef",
		AuthData:       "invalid_auth_data-request",
	}

	response, err := cardPaymentService.TokenizeCardRecurrent(request)
	assert.Error(t, err, "expected error")
	assert.Nil(t, response, "expected nil response")
}

func TestTokenizeCardRecurrentNullValues(t *testing.T) {
	gin.SetMode(gin.TestMode)
	request1 := cardPaymentServiceRequests.TokenizeCardRequest{
		TransactionRef: "",
		AuthData:       "",
	}

	response, err := cardPaymentService.TokenizeCardRecurrent(request1)
	assert.Error(t, err, "TranRef or AuthData empty")
	assert.Nil(t, response, "expected card type not empty")
}

func TestPurchaseRecurrentSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	request := cardPaymentServiceRequests.PurchaseRecurrentRequest{
		CustomerId:      "johndoe@gmail.com",
		Amount:          "5000.00",
		Currency:        "NGN",
		Token:           "512341105007817082",
		TokenExpiryDate: "5003",
		TransactionRef:  "yGDaD36ESYTXuTt",
	}

	response, err := cardPaymentService.PurchaseRecurrent(request)
	assert.NotNil(t, response)
	assert.Equal(t, response.Amount, "5000.00", err)
}

func TestPurchaseRecurrentFailure(t *testing.T) {
	gin.SetMode(gin.TestMode)

	request := cardPaymentServiceRequests.PurchaseRecurrentRequest{
		CustomerId:      "",
		Amount:          "5000",
		Currency:        "NGN",
		Token:           "512341105007817082",
		TokenExpiryDate: "5003",
		TransactionRef:  "invalid-transaction-ref",
	}

	response, err := cardPaymentService.PurchaseRecurrent(request)
	assert.NotNil(t, err, "expected an error")
	assert.Nil(t, response, "expected the response to be nil")
}

func TestGetTransactionsSuccess(t *testing.T) {
	request := cardPaymentServiceRequests.GetTransactionsRequest{
		MerchantCode:         credentialConfig.MERCHANT_CODE,
		StartDate:            "05/29/2024",
		EndDate:              "05/29/2024",
		PageSize:             1,
		PageNum:              32,
		TransactionReference: "x14NPCcOnfkOipa",
	}

	response, err := cardPaymentService.GetTransactions(request)
	assert.Nil(t, err, "no errors")
	assert.NotEmpty(t, response.Count, "count should not be empty")
}

func TestConfirmDynamicTransferSuccess(t *testing.T) {
	request := cardPaymentServiceRequests.ConfirmDynamicTransferRequest{
		MerchantCode:         credentialConfig.MERCHANT_CODE,
		TransactionReference: "ndHHBZ7PpFNIc2W",
	}

	response, err := cardPaymentService.ConfirmDynamicTransfer(request)

	assert.Nil(t, err, "expected no error")
	assert.NotNil(t, response, "expected a response")
}

func TestConfirmDynamicTransferWithInvalidReferenceOrMerchantCode(t *testing.T) {
	request := cardPaymentServiceRequests.ConfirmDynamicTransferRequest{
		MerchantCode:         "invalid_merchant_code",
		TransactionReference: "invalid-transaction-ref",
	}

	response, err := cardPaymentService.ConfirmDynamicTransfer(request)

	assert.NotNil(t, err, "expected an error")
	assert.Nil(t, response, "expected no response")
}

func TestGetRefundSuccess(t *testing.T) {
	request := cardPaymentServiceRequests.GetRefundRequest{
		MerchantCode: credentialConfig.MERCHANT_CODE,
		PageNum:      "1",
		PageSize:     "100",
		StartDate:    "05/22/2024",
		EndDate:      "05/22/2024",
	}

	response, err := cardPaymentService.GetRefund(request)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 2, response.Count)
	assert.Len(t, response.Content, 2)

	firstRefund := response.Content[0]
	assert.Equal(t, 133644, firstRefund.ID)
}

func TestGetRefundInvalidMerchantCode(t *testing.T) {
	request := cardPaymentServiceRequests.GetRefundRequest{
		MerchantCode: "INVALID_CODE",
		PageNum:      "1",
		PageSize:     "10",
		StartDate:    "2023-01-01",
		EndDate:      "2023-12-31",
	}

	response, err := cardPaymentService.GetRefund(request)

	assert.NotNil(t, err, "expected an error")
	assert.Nil(t, response, "expected no response")
	assert.Error(t, err, "expected an error")
}

func TestGetRefundInfo(t *testing.T) {
	request := cardPaymentServiceRequests.GetRefundInfoRequest{
		RefundReference: "MX6072kjsd313sjdsyegdie9322ssaud0004",
	}

	response, err := cardPaymentService.GetRefundInfo(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestGetRefundInfoInvalidReference(t *testing.T) {
	request := cardPaymentServiceRequests.GetRefundInfoRequest{
		RefundReference: "invalid_reference",
	}

	response, err := cardPaymentService.GetRefundInfo(request)
	assert.NotNil(t, err, "expected an error")
	assert.Nil(t, response, "expected no response")
	assert.Equal(t, "invalid refund reference", err.Error(), "expected error message to match")
}

func TestCreateRefundTransactionSuccess(t *testing.T) {
	randomRefundReference := credentialConfig.MERCHANT_CODE + generateRandomString(24)
	request := cardPaymentServiceRequests.CreateRefundTransactionRequest{

		RefundReference:            randomRefundReference,
		ParentTransactionReference: "QcYm7bz1N0ckTSb",
		RefundAmount:               "10",
	}

	response, err := cardPaymentService.CreateRefundTransaction(request)
	assert.Nil(t, err)
	assert.Equal(t, response.RefundReference, randomRefundReference)
	assert.Equal(t, response.MerchantCode, credentialConfig.MERCHANT_CODE)
}

func TestCreateRefundTransactionWithInvalidRequestObject(t *testing.T) {
	request := cardPaymentServiceRequests.CreateRefundTransactionRequest{

		RefundReference:            "invalid_Refund_Reference",
		ParentTransactionReference: "Invalid_Transaction_Reference",
		RefundAmount:               "10",
	}

	response, err := cardPaymentService.CreateRefundTransaction(request)
	assert.NotNil(t, err)
	assert.Nil(t, response, "expected no response")
}

func TestPayWithUSSDSuccess(t *testing.T) {
	request := cardPaymentServiceRequests.PayWithUSSDRequest{
		Amount:                       "5000",
		BankCode:                     "GTB",
		Surcharge:                    "0",
		CurrencyCode:                 "566",
		MerchantTransactionReference: "DhbRpGE1KpHmLPK",
	}

	response, err := cardPaymentService.PayWithUSSD(request)
	assert.Nil(t, err)
	assert.NotNil(t, response, "expected a response")
	assert.Equal(t, response.TransactionReference, request.MerchantTransactionReference)
}

func TestGenerateQRSuccess(t *testing.T) {
	request := cardPaymentServiceRequests.GenerateQRRequest{
		Amount:                       "5000",
		Surcharge:                    "0",
		CurrencyCode:                 "566",
		MerchantTransactionReference: "2222aaa1123h",
	}

	response, err := cardPaymentService.GenerateQR(request)
	assert.Nil(t, err)
	assert.NotNil(t, response, "expected a response")
}

func TestGenerateQRInvalidRequestObject(t *testing.T) {
	request := cardPaymentServiceRequests.GenerateQRRequest{
		MerchantTransactionReference: "invalid_Transaction_Reference",
	}

	response, err := cardPaymentService.GenerateQR(request)
	assert.NotNil(t, err)
	assert.Nil(t, response, "expected no response")
}

func TestPayWithTransferSuccess(t *testing.T) {
	request := cardPaymentServiceRequests.PayWithTransferRequest{
		MerchantCode:         "MX6072",
		PayableCode:          "9405967",
		CurrencyCode:         "566",
		Amount:               "100000",
		TransactionReference: credentialConfig.TRANSACTIONREF,
	}

	response, err := cardPaymentService.PayWithTransfer(request)
	assert.Nil(t, err)
	assert.NotNil(t, response, "expected a response")
}

func TestPayWithTransfer_invalidRequests(t *testing.T) {
	request := cardPaymentServiceRequests.PayWithTransferRequest{
		MerchantCode:         "invalid_MerchantCode",
		PayableCode:          "invalid_Payment_Code",
		TransactionReference: "invalid_Transaction_Reference",
	}

	response, err := cardPaymentService.PayWithTransfer(request)
	assert.NotNil(t, err)
	assert.Nil(t, response, "expected a response")
	assert.Error(t, err, "expected an error")
	assert.Equal(t, err.Error(), utils.FailedToPayWithTransfer)
}

func TestGetWalletCardsSuccess(t *testing.T) {
	request := cardPaymentServiceRequests.GetWalletCardsRequest{
		Username: "bughunting005@gmail.com",
		Password: "password",
	}

	response, err := cardPaymentService.GetWalletCards(request)
	assert.Nil(t, err)
	assert.NotNil(t, response, "expected a response")

	assert.NotEmpty(t, response.PaymentMethods, "expected payment methods")
	assert.Equal(t, "bughunting005@gmail.com", response.User.Email)
	assert.Equal(t, "2348081094611", response.User.MobileNo)

	for _, method := range response.PaymentMethods {
		assert.NotEmpty(t, method.CardTypeCode, "expected card type code")
		assert.NotEmpty(t, method.MaskedPan, "expected masked PAN")
		assert.NotEmpty(t, method.CardIdentifier, "expected card identifier")
	}
}

func TestGetWalletCards_invalidCredential(t *testing.T) {
	request := cardPaymentServiceRequests.GetWalletCardsRequest{
		Username: "invalid_username",
		Password: "invalid_password",
	}

	response, err := cardPaymentService.GetWalletCards(request)
	assert.NotNil(t, err)
	assert.Error(t, err, "expected an error")
	assert.Nil(t, response, "expected no response")
}

func TestGenerateAlternativePaymentOptionsSuccess(t *testing.T) {
	response, err := cardPaymentService.GenerateAlternativePaymentOptions()
	assert.Nil(t, err)
	assert.NotNil(t, response, "expected a response")

	assert.Equal(t, "sadd", response.Merchant.Name)
	assert.Equal(t, "MX6072", response.Merchant.MerchantCode)
	assert.NotEmpty(t, response.PaymentOptions, "expected payment options")

	for _, option := range response.PaymentOptions {
		assert.NotEmpty(t, option.PayableCode, "expected payable code")
		assert.NotEmpty(t, option.ProviderCode, "expected provider code")
		assert.NotNil(t, option.Enabled, "expected enabled status")
		fmt.Println(option.AdditionalInformation)
	}
}

func TestGetUSSDBanksSuccess(t *testing.T) {
	banks, err := cardPaymentService.GetUSSDBanks()
	assert.Nil(t, err)
	assert.NotNil(t, banks, "expected a list of banks")
}
