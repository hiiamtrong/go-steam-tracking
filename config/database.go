package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func newMongoClient() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/steam_tracking"))

	if err != nil {
		log.Fatalln(err)
	}

	DB = client
	log.Println("MongoDB connected")
}
