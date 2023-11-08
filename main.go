package main

import (
	"taksa/bot"

	"github.com/davecgh/go-spew/spew"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	err := <-bot.Init()
	if err != nil {
		panic(err)
	}

	botInstance := bot.GetBot()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := botInstance.GetUpdatesChan(u)

	for update := range updates {
		var msg tgbotapi.MessageConfig
		// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		spew.Dump(update)

		if update.Message != nil {

			if update.Message.Chat.Type == "group" {

				if update.Message.NewChatMembers != nil && update.Message.NewChatMembers[0].ID == botInstance.Self.ID { // when bot was added to chat
					if ok := bot.AddToChat(update.Message.Chat.ID); !ok {
						// handle an error
					}
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Taksa bot is here")
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "test")
				}
				msg.ReplyToMessageID = update.Message.MessageID
				botInstance.Send(msg)
			} else {
				msg = tgbotapi.NewMessage(update.Message.From.ID, "test")
				if update.Message.IsCommand() {
					switch update.Message.Command() {
					case "start": // when user starts a new private chat with bot
						if ok := bot.StartPrivateChat(update.Message.From.ID); !ok {
							// handle an error
						}
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Hello you")

					default:
						continue
					}

				} else {

				}
				botInstance.Send(msg)
			}
		}

	}
}
