package main

import (
	"log"
	"taksa/bot"
	"taksa/db"
	"taksa/texts"

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

	for update := range updates {
		var msg tgbotapi.MessageConfig

		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.Chat.Type == "group" {
				if update.Message.IsCommand() {
					switch update.Message.Command() {
					case "new_order":
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "catch!")
					default:
						continue
					}
				} else if update.Message.NewChatMembers != nil && update.Message.NewChatMembers[0].ID == botInstance.Self.ID { // when bot was added to chat
					ok, err := bot.AddToChat(update.Message.Chat)
					if err != nil {
						panic(err)
					}
					if ok {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, texts.MessageAddToGroup)
					} else {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, texts.MessageAlreadyAddedToGroup)
					}
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
						isNewUser, err := bot.StartPrivateChat(update.Message.From)
						if err != nil {
							// handle an error
						}
						if isNewUser {
							msg = tgbotapi.NewMessage(update.Message.Chat.ID, texts.MessageWelcomePrivateChat)
						} else {
							msg = tgbotapi.NewMessage(update.Message.Chat.ID, texts.MessageWelcomePrivateChatWelcomeBack)
						}
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
