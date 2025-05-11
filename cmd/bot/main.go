package main

import (
	"fmt"
	"log"

	"github.com/TimNikolaev/Pocketoid/configs"
	"github.com/TimNikolaev/Pocketoid/internal/repository/boltdb"
	"github.com/TimNikolaev/Pocketoid/internal/server"
	"github.com/TimNikolaev/Pocketoid/internal/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
)

func main() {
	config, err := configs.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config)

	bot, err := tgbotapi.NewBotAPI(config.TgToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient(config.ConsumerKey)
	if err != nil {
		log.Fatal(err)
	}

	boltRepository, err := boltdb.NewRepository(config.BDPath)
	if err != nil {
		log.Fatal(err)
	}

	tgBot := telegram.NewBot(bot, pocketClient, boltRepository, config.AuthServerURL)

	authServer := server.NewAuthorizationServer(pocketClient, boltRepository, config.TgBotURL)

	go func() {
		if err := tgBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err = authServer.Start(); err != nil {
		log.Fatal(err)
	}
}
