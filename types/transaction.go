package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`
	Name      string               `bson:"name"`
	Amount    int32                `bson:"amount"`
	SplitType string               `bson:"split_type"`
	PaidBy    primitive.ObjectID   `bson:"paid_by"`
	Splits    []primitive.ObjectID `bson:"splits"`
	Events    []primitive.ObjectID `bson:"events"`
}
