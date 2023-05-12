package config

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

type Config struct {
	SnipcartApiKey string `env:"SNIPCART_API_KEY,unset"`
	AwsSmsArn      string `env:"API_SMS_SECRET_ARN,unset"`
	Production     bool   `env:"API_PRODUCTION"`
}

type WebhookSmsSecret struct {
	SnipcartApiKey string `json:"snipcart_api_key"`
}

func NewConfigFromFile(filePath string) (*Config, error) {
	if err := godotenv.Load(filePath); err != nil {
		return &Config{}, err
	}

	// loading from a file so we assume not to use Sms, maybe change this in the
	// future?
	config, err := NewConfigFromEnv(false)
	if err != nil {
		return config, err
	}

	return config, nil
}

func NewConfigFromEnv(useAwsSms bool) (*Config, error) {
	var config Config
	if err := env.Parse(&config); err != nil {
		return &Config{}, err
	}

	if useAwsSms {
		secretCache, err := secretcache.New()
		if err != nil {
			return &config, nil
		}

		var webhookSmsSecret WebhookSmsSecret
		secretString, err := secretCache.GetSecretString(config.AwsSmsArn)
		if err != nil {
			return &config, fmt.Errorf("issue with GetSecretString: %s\n", err.Error())
		}

		err = json.Unmarshal([]byte(secretString), &webhookSmsSecret)
		if err != nil {
			return &config, fmt.Errorf("issue with unmarshal: %s\n", err.Error())
		}

		config.SnipcartApiKey = webhookSmsSecret.SnipcartApiKey
	}

	return &config, nil
}
