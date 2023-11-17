package bot

import (
	"context"
	"taksa/db"
	"taksa/types"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson"
)

func StartPrivateChat(user *tgbotapi.User) (bool, error) {
	/*
	* @Todo:
	*
	* Реализовать запись в базу,
	* обработчики на success и error от базы
	* написаны ниже
	*
	 */

	collection := db.GetCollection("users")

	filter := bson.M{"username": user.UserName}
	var existingParticipant types.Participant

	err := collection.FindOne(context.TODO(), filter).Decode(&existingParticipant)
	if err == nil {
		return false, nil
	}

	newChat := types.Participant{
		Username: user.UserName,
		Fullname: user.FirstName + " " + user.LastName,
	}
	// Insert user document into MongoDB
	_, err = collection.InsertOne(context.TODO(), newChat)
	if err != nil {
		return false, err
	}

	return true, nil
}
