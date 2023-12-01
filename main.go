package main

import (
	_"log"
	"taksa/bot"
	"taksa/db"
	_"taksa/texts"
	"taksa/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	db.Init()
	err := <-bot.Init()
	if err != nil {
		panic(err)
	}

	botInstance := bot.GetBot()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := botInstance.GetUpdatesChan(u)
	handlers.HandleUpdate(updates, botInstance)

}
