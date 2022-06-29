package main

import (
	"context"
	"time"

	"github.com/hiiamtrong/golang-steam-tracking/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func main() {

	config.Setup()

	coll := config.DB.Database("steam_tracking").Collection("game_details")
	index := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "name", Value: bsonx.String("text")}},
		},
	}

	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	_, errIndex := coll.Indexes().CreateMany(context.TODO(), index, opts)
	if errIndex != nil {
		panic(errIndex)
	}
}
