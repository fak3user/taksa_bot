package bot

import (
	"context"
	"taksa/db"
	"taksa/types"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson"
)

func AddToChat(chat *tgbotapi.Chat) (bool, error) {
	/*
	* @Todo:
	*
	* Реализовать запись в базу,
	* обработчики на success и error от базы
	* написаны ниже
	*
	 */

	collection := db.GetCollection("chats")

	filter := bson.M{"tg_id": chat.ID}
	var existingChat types.Chat

	err := collection.FindOne(context.TODO(), filter).Decode(&existingChat)
	if err == nil {
		return false, nil
	}

	newChat := types.Chat{
		TgID: chat.ID,
		Name: chat.Title,
	}
	// Insert user document into MongoDB
	_, err = collection.InsertOne(context.TODO(), newChat)
	if err != nil {
		return false, err
	}

	return true, nil
}

func RemoveFromChat(chatId int64) bool {
	/*
	* @Todo:
	*
	* Реализовать запись в базу,
	* обработчики на success и error от базы
	* написаны ниже
	*
	 */

	if true { // onSuccess
		return true
	} else { // onError
		return false
	}
}
