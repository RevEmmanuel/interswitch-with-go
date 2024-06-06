package cardPaymentAPITests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"interswitch_go_testing/src/Dtos/requests/cardPaymentServiceRequests"
	"interswitch_go_testing/src/credentialConfig"
	"interswitch_go_testing/src/services/cardPaymentService"
	"log"
	"testing"
)

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
		fmt.Println("Balance: ", response.Balance)
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
	fmt.Println(response)
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

	if err != nil {
		log.Printf("Received expected error: %v", err)
	} else {
		log.Fatal("Expected error but got none")
	}
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
	fmt.Printf("Response: %+v\n", response)
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
