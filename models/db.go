package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id            bson.ObjectId  `json:"id"            bson:"_id"`
	Name          string         `json:"name"          bson:"name"`
	Email         string         `json:"email"         bson:"email"`
	Role          string         `json:"role"          bson:"role"`
	Subscriptions []Subscription `json:"subscriptions" bson:"subscriptions"`
}

type Subscription struct {
	Name  string  `json:"name"  bson:"name"`
	Price float64 `json:"price" bson:"price"`
}
