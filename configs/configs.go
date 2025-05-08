package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TgToken     string
	ConsumerKey string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		TgToken:     os.Getenv("TG_TOKEN"),
		ConsumerKey: os.Getenv("CONSUMER_KEY"),
	}, nil
}
