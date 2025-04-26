package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var DB *mongo.Database

func Config() {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_ = client.Ping(ctx, readpref.Primary())

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	DB = client.Database("SocialChef")

}
