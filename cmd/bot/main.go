package main

import (
	"log"

	"github.com/TimNikolaev/Pocketoid/configs"
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

	pocketClient, err := pocket.NewClient(config.ConsumerKey)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	tgBot := telegram.NewBot(bot, pocketClient, "http://localhost/")

	if err := tgBot.Start(); err != nil {
		log.Fatal(err)
	}
}
