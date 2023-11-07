package main

import (
	"fmt"
	"log"
	"taksa/bot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	err := <-bot.Init()
	if err != nil {
		panic(err)
	}

	tgbot := bot.GetBot()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tgbot.GetUpdatesChan(u)

	for update := range updates {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		echoMsg := tgbotapi.NewMessage(update.Message.From.ID, update.Message.Text)
		tgbot.Send(echoMsg)
	}
	fmt.Println("321")
}
