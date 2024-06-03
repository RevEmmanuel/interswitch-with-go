package authService

import (
	"errors"
	"github.com/go-co-op/gocron"
	"github.com/go-resty/resty/v2"
	"interswitch_go_testing/src/config"
	"log"
	"time"
)

var token string

type InterswitchAuthenticationResponse struct {
	AccessToken           string   `json:"access_token"`
	TokenType             string   `json:"token_type"`
	Bearer                string   `json:"bearer"`
	ExpiresIn             int64    `json:"expires_in"`
	Scope                 string   `json:"scope"`
	MerchantCode          string   `json:"merchant_code"`
	ProductionPaymentCode string   `json:"production_payment_code"`
	ProductId             int64    `json:"productId"`
	RequestorId           string   `json:"requestor_id"`
	Env                   string   `json:"env"`
	PayableId             string   `json:"payable_id"`
	ClientDescription     string   `json:"client_description"`
	InstitutionId         string   `json:"institution_id"`
	CoreId                string   `json:"core_id"`
	ApiResources          []string `json:"api_resources"`
	PostcardUser          string   `json:"postcard_user"`
	ClientName            string   `json:"client_name"`
	ClientLogo            string   `json:"client_logo"`
	PaymentCode           string   `json:"payment_code"`
	Jti                   string   `json:"jti"`
}

func Init() error {
	// Perform initialization here (e.g., renew authentication token)
	err := renewAuthenticationToken()
	if err != nil {
		return err
	}
	return nil
}

func renewAuthenticationToken() error {
	client := resty.New()
	formData := map[string]string{
		"grant_type": "client_credentials",
		"scope":      "profile",
	}

	resp, err := client.R().
		//EnableTrace().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("Authorization", config.Config.AuthHeader).
		SetFormData(formData).
		SetResult(&InterswitchAuthenticationResponse{}).
		Post(config.Config.AuthURL)

	if err != nil {
		log.Printf("Failed to renew token: %v", err)
		return err
	}

	if resp.StatusCode() != 200 {
		log.Printf("Failed to renew token: received status code %d", resp.StatusCode())
		return errors.New("failed to renew token: non-200 status code received")
	}

	response := resp.Result().(*InterswitchAuthenticationResponse)
	token = response.AccessToken
	log.Printf("Token renewed: %s", token)
	return nil
}

func StartTokenRenewal() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().At("00:00").Do(func() {
		if err := renewAuthenticationToken(); err != nil {
			log.Printf("Error during scheduled token renewal: %v", err)
		}
	})
	s.StartAsync()
}

func GetToken() string {
	return token
}
