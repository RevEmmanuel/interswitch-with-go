package cardPaymentService

import (
	"bytes"
	"encoding/json"
	"errors"
	"interswitch_go_testing/src/Dtos/requests/cardPaymentServiceRequests"
	"interswitch_go_testing/src/Dtos/responses"
	"interswitch_go_testing/src/credentialConfig"
	"interswitch_go_testing/src/utils"
	"io"
	"net/http"
	"time"
)

func TokenizeCardRecurrent(request cardPaymentServiceRequests.TokenizeCardRequest) (*responses.TokenizeCardResponse, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		credentialConfig.TOKENIZETRANSACTIONURL,
		bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", credentialConfig.AUTHTOKEN)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response responses.TokenizeCardResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return nil, err
		}
		return &response, nil
	}

	return nil, errors.New(utils.UnableToTokenizeCard)
}

func PurchaseRecurrent(request cardPaymentServiceRequests.PurchaseRecurrentRequest) (*responses.PurchaseRecurrentResponse, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		credentialConfig.VALIDATE_PURCHASE_RECURRENT_URL,
		bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", credentialConfig.AUTHTOKEN)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		if cerr := Body.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}(resp.Body)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response responses.PurchaseRecurrentResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return nil, err
		}
		return &response, nil
	}

	return nil, errors.New(utils.FailedToProcessRecurrentPurchase)
}
