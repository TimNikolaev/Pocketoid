package main

import (
	"log"

	"github.com/TimNikolaev/Pocketoid/configs"
	bolt_repository "github.com/TimNikolaev/Pocketoid/pkg/repository/bolt"
	"github.com/TimNikolaev/Pocketoid/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
)

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(config.TgToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient(config.ConsumerKey)
	if err != nil {
		log.Fatal(err)
	}

	boltRepository, err := bolt_repository.NewRepository()
	if err != nil {
		log.Fatal(err)
	}

	tgBot := telegram.NewBot(bot, pocketClient, boltRepository, "http://localhost/")

	if err := tgBot.Start(); err != nil {
		log.Fatal(err)
	}
}
