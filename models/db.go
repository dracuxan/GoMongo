package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id            primitive.ObjectID `json:"id"            bson:"_id"`
	Name          string             `json:"name"          bson:"name"`
	Email         string             `json:"email"         bson:"email"`
	Role          string             `json:"role"          bson:"role"`
	Subscriptions []Subscription     `json:"subscriptions" bson:"subscriptions"`
}

type Subscription struct {
	Name  string  `json:"name"  bson:"name"`
	Price float64 `json:"price" bson:"price"`
}
