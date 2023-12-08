package handlers

import (
	"log"
	"taksa/accounts"
	"taksa/bot"
	"taksa/texts"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdate(updates tgbotapi.UpdatesChannel, botInstance *tgbotapi.BotAPI) {
	for update := range updates {
		if update.Message != nil {
			logMessage(update)
			handleMessage(update, botInstance)
		}
	}
}

// logMessage logs the incoming message
func logMessage(update tgbotapi.Update) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
}

// handleMessage processes the received message based on its type and content
func handleMessage(update tgbotapi.Update, botInstance *tgbotapi.BotAPI) {
	var msg tgbotapi.MessageConfig

	if update.Message.Chat.Type == "group" {
		msg = handleGroupMessage(update, botInstance)
	} else {
		msg = handlePrivateMessage(update, botInstance)
	}

	if msg.Text != "" {
		msg.ReplyToMessageID = update.Message.MessageID
		botInstance.Send(msg)
	}
}

// handleGroupMessage processes messages from group chats
func handleGroupMessage(update tgbotapi.Update, botInstance *tgbotapi.BotAPI) tgbotapi.MessageConfig {
	if update.Message.IsCommand() {
		return handleGroupCommand(update)
	} else if isNewMember(update, botInstance) {
		return handleNewGroupMember(update)
	}
	return tgbotapi.NewMessage(update.Message.Chat.ID, "test") // default response
}

// handlePrivateMessage processes messages from private chats
func handlePrivateMessage(update tgbotapi.Update, botInstance *tgbotapi.BotAPI) tgbotapi.MessageConfig {
	if update.Message.IsCommand() {
		return handlePrivateCommand(update, botInstance)
	}
	return tgbotapi.NewMessage(update.Message.From.ID, "test") // default response
}

func handleGroupCommand(update tgbotapi.Update) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig
	switch update.Message.Command() {
	case "new_order":
		account, err := accounts.CreateAccount()
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Catch!")
		// Add more cases as necessary
	}
	return msg
}

func handleNewGroupMember(update tgbotapi.Update) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig
	ok, err := bot.AddToChat(update.Message.Chat)
	if err != nil {
		panic(err)
	}
	if ok {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, texts.MessageAddToGroup)
	} else {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, texts.MessageAlreadyAddedToGroup)
	}
	return msg
}

func handlePrivateCommand(update tgbotapi.Update, botInstance *tgbotapi.BotAPI) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig
	switch update.Message.Command() {
	case "start":
		// Implement logic for start command
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome!")
		// Add more cases as necessary
	}
	return msg
}

func isNewMember(update tgbotapi.Update, botInstance *tgbotapi.BotAPI) bool {
	// Assuming the first user in NewChatMembers is the one who joined
	return update.Message.NewChatMembers != nil && len(update.Message.NewChatMembers) > 0 && update.Message.NewChatMembers[0].ID == botInstance.Self.ID
}
