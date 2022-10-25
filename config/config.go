package config

import (
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"log"
)

type Config struct {
	HTTP  HTTPConfig
	OAUTH oauth2.Config
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		HTTP:  LoadHTTPConfig(),
		OAUTH: LoadOAUTHConfiguration(),
	}
}
