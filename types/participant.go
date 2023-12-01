package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Participant struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty"`
	Username      string               `bson:"username"`
	Fullname      string               `bson:"fullname"`
	Notifications []primitive.ObjectID `bson:"notifications"`
	Accounts        []primitive.ObjectID `bson:"Accounts"`
}
