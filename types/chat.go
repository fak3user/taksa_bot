package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chat struct {
	ID     primitive.ObjectID   `bson:"_id,omitempty"`
	TgID   int64                `bson:"tg_id"`
	Name   string               `bson:"name"`
	Accounts []primitive.ObjectID `bson:"Accounts"`
}
