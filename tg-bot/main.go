package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"tg-bot/handler"
)

const (
	ConsulAddr = "127.0.0.1:8500"
)

var (
	service = "tg-bot"
	host    = "127.0.0.1"
	port    = ":12005"
	version = "latest"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("6680679731:AAHLFR2QhioohAo7rwqCC_2azsYr7JdHKYE")
	if err != nil {
		// Abort if something is wrong
		logger.Log(logger.ErrorLevel, "bot err: ", err)
	}

	bot.Debug = false
	// Create service
	srv := micro.NewService(
		micro.Address(port),
		micro.Name(service),
		micro.Handle(handler.TgBot{Bot: bot}),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}

}
