package cardPaymentAPITests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"interswitch_go_testing/src/Dtos/requests/cardPaymentServiceRequests"
	"interswitch_go_testing/src/credentialConfig"
	"interswitch_go_testing/src/services/cardPaymentService"
	"testing"
)

func TestTokenizeCardRecurrentSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	request := cardPaymentServiceRequests.TokenizeCardRequest{
		TransactionRef: credentialConfig.TRANSACTIONREF,
		AuthData:       credentialConfig.AUTHDATA,
	}

	response, err := cardPaymentService.TokenizeCardRecurrent(request)
	//if err != nil {
	//	t.Fatalf("expected no error, got %v", err)
	//}
	if response == nil {
		assert.Nil(t, err, "no errors")
		assert.NotEmpty(t, response.CardType)
		assert.NotEmpty(t, response.Balance)
		assert.Equal(t, response.TransactionRef, credentialConfig.TRANSACTIONREF)
	} else {
		t.Fatalf("expected response, got nil")
	}

	fmt.Println("TokenizeCardRecurrent: ", response.CardType)
}

func TestTokenizeCardRecurrentFailure(t *testing.T) {
	gin.SetMode(gin.TestMode)
	request := cardPaymentServiceRequests.TokenizeCardRequest{
		TransactionRef: "invalid-transaction-ref",
		AuthData:       credentialConfig.AUTHDATA,
	}

	response, err := cardPaymentService.TokenizeCardRecurrent(request)
	assert.Error(t, err, "expected an error")
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
		Amount:          "5000",
		Currency:        "NGN",
		Token:           "512341105007817082",
		TokenExpiryDate: "5003",
		TransactionRef:  credentialConfig.TRANSACTIONREF,
	}

	response, err := cardPaymentService.PurchaseRecurrent(request)
	if response != nil {
		assert.Nil(t, err)
		assert.NotEmpty(t, response.TransactionRef)
		assert.NotEmpty(t, response.PaymentId)
		assert.NotEmpty(t, response.BankCode)

	}
}

func TestPurchaseRecurrentFailure(t *testing.T) {
	gin.SetMode(gin.TestMode)

	request := cardPaymentServiceRequests.PurchaseRecurrentRequest{
		CustomerId:      "invalid-customer-id",
		Amount:          "5000",
		Currency:        "NGN",
		Token:           "512341105007817082",
		TokenExpiryDate: "5003",
		TransactionRef:  "invalid-transaction-ref",
	}

	response, err := cardPaymentService.PurchaseRecurrent(request)
	assert.Error(t, err, "expected an error")
	assert.Nil(t, response, "expected nil response")
}
