package main

import "go.mongodb.org/mongo-driver/v2/bson"

type Item struct {
	Name     string `bson:"name" json:"name"`
	Quantity int    `bson:"quantity" json:"quantity"`
}

type Order struct {
	ID           bson.ObjectID `bson:"_id,omitempty" json:"id"`
	CustomerName string             `bson:"customerName" json:"customName"`
	Phone        string             `bson:"phone" json:"phone"`
	Items        []Item             `bson:"items" json:"items"`
	Status       string             `bson:"status" json:"status"`
	Platform     string             `bson:"platform" json:"platform"`
	Timestamp    string             `bson:"timestamp" json:"timestamp"`
}

type User struct {
	ID       bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
}

type UserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
