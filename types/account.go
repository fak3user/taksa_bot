package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty"`
	Name         string               `bson:"name"`
	Description  string               `bson:"description"`
	Opened       bool                 `bson:"opened"`
	Transactions []primitive.ObjectID `bson:"transactions"`
	Participants []primitive.ObjectID `bson:"participants"`
}
