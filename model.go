package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Username    string             `bson:"username,omitempty"`
	PhoneNumber string             `bson:"phone_number,omitempty"`
}

type Invoice struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserId      primitive.ObjectID `bson:"user_id,omitempty"`
	Amount      int                `bson:"amount,omitempty"`
	ChannelType string             `bson:"channel_type,omitempty"`
	Status      string             `bson:"status,omitempty"`
}

type Payment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	InvoiceId primitive.ObjectID `bson:"invoice_id,omitempty" json:"invoice_id"`
	Amount    int                `bson:"amount,omitempty" json:"amount"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
}
