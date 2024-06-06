package cardPaymentService

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"interswitch_go_testing/src/Dtos/requests/cardPaymentServiceRequests"
	"interswitch_go_testing/src/Dtos/responses"
	"interswitch_go_testing/src/credentialConfig"
	"interswitch_go_testing/src/utils"
)

func TokenizeCardRecurrent(request cardPaymentServiceRequests.TokenizeCardRequest) (*responses.TokenizeCardResponse, error) {

	client := resty.New()
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", credentialConfig.AUTHTOKEN).
		SetBody(requestBody).
		SetResult(&responses.TokenizeCardResponse{}).
		Post(credentialConfig.TOKENIZETRANSACTIONURL)

	if resp.IsError() {
		//return nil, errors.New("request failed with status " + resp.Status)
		return nil, errors.New(utils.UnableToTokenizeCard)
	}
	if err != nil {
		return nil, err
	}

	response := resp.Result().(*responses.TokenizeCardResponse)
	return response, nil

}

func PurchaseRecurrent(request cardPaymentServiceRequests.PurchaseRecurrentRequest) (*responses.PurchaseRecurrentResponse, error) {
	client := resty.New()

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", credentialConfig.AUTHTOKEN).
		SetBody(requestBody).
		SetResult(&responses.PurchaseRecurrentResponse{}).
		Post(credentialConfig.VALIDATE_PURCHASE_RECURRENT_URL)

	if resp.IsError() {
		return nil, errors.New(utils.FailedToProcessRecurrentPurchase)
	}
	if err != nil {
		return nil, err
	}

	response := resp.Result().(*responses.PurchaseRecurrentResponse)
	return response, nil
}

func GetTransactions(request cardPaymentServiceRequests.GetTransactionsRequest) (*responses.GetTransactionsResponse, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", credentialConfig.AUTHTOKEN).
		SetQueryParams(map[string]string{
			"merchantCode":         request.MerchantCode,
			"transactionReference": request.TransactionReference,
			"startDate":            request.StartDate,
			"endDate":              request.EndDate,
			"pageSize":             fmt.Sprintf("%d", request.PageSize),
			"pageNum":              fmt.Sprintf("%d", request.PageNum),
		}).
		SetResult(&responses.GetTransactionsResponse{}).
		Get(credentialConfig.GET_TRANSACTION_URL)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, errors.New(utils.FailedToGetTransactions)
	}

	response := resp.Result().(*responses.GetTransactionsResponse)
	return response, nil
}

func ConfirmDynamicTransfer(request cardPaymentServiceRequests.ConfirmDynamicTransferRequest) (*responses.ConfirmDynamicTransferResponse, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", credentialConfig.AUTHTOKEN).
		SetQueryParams(map[string]string{
			"merchantCode":         request.MerchantCode,
			"transactionReference": request.TransactionReference,
		}).
		SetResult(&responses.ConfirmDynamicTransferResponse{}).
		Get(credentialConfig.CONFIRM_DYNAMIC_TRANSFER_URL)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, errors.New(utils.FailedToConfirmDynamicTransfer)
	}

	response := resp.Result().(*responses.ConfirmDynamicTransferResponse)
	return response, nil
}

func GetRefund(request cardPaymentServiceRequests.GetRefundRequest) (*responses.GetRefundResponse, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", credentialConfig.AUTHTOKEN).
		SetQueryParams(map[string]string{
			"merchantCode": request.MerchantCode,
			"pageNum":      request.PageNum,
			"pageSize":     request.PageSize,
			"startDate":    request.StartDate,
			"endDate":      request.EndDate,
		}).
		SetResult(&responses.GetRefundResponse{}).
		Get(credentialConfig.GET_REFUND_URL)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, errors.New(utils.FailedToGetRefund)
	}

	response := resp.Result().(*responses.GetRefundResponse)
	return response, nil
}
