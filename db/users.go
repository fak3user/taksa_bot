package db

import (
	"context"
	"taksa/types"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson"
)

func AddUserAndCheckIfExist(user *tgbotapi.User) (types.Participant, bool, error) {
	collection := GetCollection("users")

	filter := bson.M{"username": user.UserName}
	var existingParticipant types.Participant

	err := collection.FindOne(context.TODO(), filter).Decode(&existingParticipant)
	if err == nil {
		return existingParticipant, false, nil
	}

	newParticipant := types.Participant{
		Username: user.UserName,
		Fullname: user.FirstName + " " + user.LastName,
	}

	// Insert user document into MongoDB
	insertResult, err := collection.InsertOne(context.TODO(), newParticipant)
	if err != nil {
		return existingParticipant, false, err
	}

	filter = bson.M{"_id": insertResult.InsertedID}
	var insertedParticipant types.Participant

	collection.FindOne(context.TODO(), filter).Decode(&insertedParticipant)

	return insertedParticipant, true, nil
}
