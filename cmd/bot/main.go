package main

import (
	"log"

	"github.com/TimNikolaev/Pocketoid/internal/config"
	"github.com/TimNikolaev/Pocketoid/internal/repository/boltdb"
	"github.com/TimNikolaev/Pocketoid/internal/server"
	"github.com/TimNikolaev/Pocketoid/internal/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TgToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient(cfg.ConsumerKey)
	if err != nil {
		log.Fatal(err)
	}

	boltRepository, err := boltdb.NewRepository(cfg.BDPath)
	if err != nil {
		log.Fatal(err)
	}

	tgBot := telegram.NewBot(bot, pocketClient, boltRepository, cfg.AuthServerURL, cfg.Messages)

	authServer := server.NewAuthorizationServer(pocketClient, boltRepository, cfg.TgBotURL)

	go func() {
		if err := tgBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err = authServer.Start(); err != nil {
		log.Fatal(err)
	}
}
