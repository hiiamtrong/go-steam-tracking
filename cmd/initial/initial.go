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

	var boolean = true

	coll := config.DB.Database("steam_tracking").Collection("game_details")
	index := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "name", Value: bsonx.String("text")}},
		},
		{
			Keys: bsonx.Doc{{Key: "steam_appid", Value: bsonx.Int32(-1)}},
			Options: &options.IndexOptions{
				Unique: &boolean,
			},
		},
	}

	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	_, errIndex := coll.Indexes().CreateMany(context.TODO(), index, opts)
	if errIndex != nil {
		panic(errIndex)
	}
}
