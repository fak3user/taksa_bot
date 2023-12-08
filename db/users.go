package db

import (
	"context"
	"taksa/types"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson"
)

func AddUserAndCheckIfExist(user *tgbotapi.User) (types.User, bool, error) {
	collection := GetCollection("users")

	filter := bson.M{"username": user.UserName}
	var existingUser types.User

	err := collection.FindOne(context.TODO(), filter).Decode(&existingUser)
	if err == nil {
		return existingUser, false, nil
	}

	newUser := types.User{
		Username: user.UserName,
		Fullname: user.FirstName + " " + user.LastName,
	}

	// Insert user document into MongoDB
	insertResult, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		return existingUser, false, err
	}

	filter = bson.M{"_id": insertResult.InsertedID}
	var insertedUser types.User

	collection.FindOne(context.TODO(), filter).Decode(&insertedUser)

	return insertedUser, true, nil
}
