package config

import (
	"github.com/joho/godotenv"
	"log"
)

type Configuration struct {
	AuthURL    string
	AuthHeader string
}

var Config Configuration

func LoadConfig() {
	config, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}

	Config = Configuration{
		AuthURL:    config["authURL"],
		AuthHeader: config["authHeader"],
	}

}
