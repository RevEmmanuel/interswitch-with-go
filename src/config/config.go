package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Configuration struct {
	AuthURL                      string
	AuthHeader                   string
	ValidatePurchaseRecurrentUrl string
	AUTHTOKEN                    string
	TOKENIZETRANSACTIONURL       string
	AUTHDATA                     string
	TRANSACTIONREF               string
}

var Config Configuration

func LoadConfig() {
	_, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}

	Config = Configuration{
		AuthURL:                      os.Getenv("AUTHURL"),
		AuthHeader:                   os.Getenv("AUTHHEADER"),
		ValidatePurchaseRecurrentUrl: os.Getenv("VALIDATE_PURCHASE_RECURRENT_URL"),
		AUTHTOKEN:                    os.Getenv("AUTHTOKEN"),
		TOKENIZETRANSACTIONURL:       os.Getenv("TOKENIZETRANSACTIONURL"),
		AUTHDATA:                     os.Getenv("AUTHDATA"),
		TRANSACTIONREF:               os.Getenv("TRANSACTIONREF"),
	}

}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
