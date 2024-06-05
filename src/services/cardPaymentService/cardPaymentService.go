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
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

func TokenizeCardRecurrent(request cardPaymentServiceRequests.TokenizeCardRequest) (*responses.TokenizeCardResponse, error) {

	client := &http.Client{Timeout: 100 * time.Second}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", credentialConfig.TOKENIZETRANSACTIONURL, bytes.NewBuffer(requestBody))
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		//return nil, errors.New("request failed with status " + resp.Status)
		return nil, errors.New(utils.UnableToTokenizeCard)
	}

	var tokenizeCardResponse responses.TokenizeCardResponse
	err = json.Unmarshal(body, &tokenizeCardResponse)
	if err != nil {
		return nil, err
	}

	return &tokenizeCardResponse, nil
}

func PurchaseRecurrent(request cardPaymentServiceRequests.PurchaseRecurrentRequest) (*responses.PurchaseRecurrentResponse, error) {

	client := &http.Client{Timeout: 100 * time.Second}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", credentialConfig.VALIDATE_PURCHASE_RECURRENT_URL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", credentialConfig.AUTHTOKEN)

	resp, err := client.Do(req)
	if err != nil {
		if urlErr, ok := err.(*url.Error); ok {
			if opErr, ok := urlErr.Err.(*net.OpError); ok {
				log.Printf("Network error: %v", opErr)
			}
		}
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusAccepted {
		return nil, errors.New(utils.FailedToProcessRecurrentPurchase)
	}

	var purchaseRecurrentResponse responses.PurchaseRecurrentResponse
	err = json.Unmarshal(body, &purchaseRecurrentResponse)
	if err != nil {
		return nil, err
	}

	return &purchaseRecurrentResponse, nil
}
