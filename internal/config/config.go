package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	TgToken       string
	ConsumerKey   string
	AuthServerURL string
	TgBotURL      string `mapstructure:"bot_url"`
	BDPath        string `mapstructure:"db_file"`

	Messages
}

type Messages struct {
	Responses
	Errors
}

type Responses struct {
	Start             string `mapstructure:"start"`
	AlreadyAuthorized string `mapstructure:"already_authorized"`
	SuccessSave       string `mapstructure:"success_save"`
	UnknownCommand    string `mapstructure:"unknown_command"`
}

type Errors struct {
	Default      string `mapstructure:"default"`
	InvalidURL   string `mapstructure:"invalid_url"`
	Unauthorized string `mapstructure:"unauthorized"`
	FailToSave   string `mapstructure:"fail_to_save"`
}

func InitConfig() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.responses", &cfg.Messages.Responses); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.errors", &cfg.Messages.Errors); err != nil {
		return nil, err
	}

	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseEnv(cfg *Config) error {
	if err := viper.BindEnv("tg_token"); err != nil {
		return err
	}

	if err := viper.BindEnv("consumer_key"); err != nil {
		return err
	}

	if err := viper.BindEnv("auth_server_url"); err != nil {
		return err
	}

	cfg.TgToken = viper.GetString("tg_token")
	cfg.ConsumerKey = viper.GetString("consumer_key")
	cfg.AuthServerURL = viper.GetString("auth_server_url")

	return nil
}

func loadEnv(cfg *Config) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	cfg.TgToken = os.Getenv("TG_TOKEN")
	cfg.ConsumerKey = os.Getenv("CONSUMER_KEY")
	cfg.AuthServerURL = os.Getenv("AUTH_SERVER_URL")

	return nil
}
