package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	Token         string
	WebhookDomain string
	WebhookSecret string
	ApiKey       string
}

func ReadEnv() (EnvConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := EnvConfig{
		Token:         os.Getenv("TOKEN"),
		WebhookDomain: os.Getenv("WEBHOOK_DOMAIN"),
		WebhookSecret: os.Getenv("WEBHOOK_SECRET"),
		ApiKey: os.Getenv("API_KEY"),
	}

	if config.Token == ""{
		panic("Token is empty")
	}

	if config.WebhookDomain == ""{
		panic("TokWebhookDomainen is empty")
	}

	if config.WebhookSecret == ""{
		panic("WebhookSecret is empty")
	}

	if config.ApiKey == ""{
		panic("ApiKey is empty")
	}

	return config, nil
}
