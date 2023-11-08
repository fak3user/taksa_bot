package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Split struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Amount      int32              `bson:"amount,omitempty"`
	Percentage  int16              `bson:"percentage,omitempty"`
	Participant primitive.ObjectID `bson:"participant"`
}
