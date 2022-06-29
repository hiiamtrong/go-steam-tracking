package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func newMongoClient() {

	strConnection := fmt.Sprintf("mongodb://%s:%s", Cfg.MONGO_HOST, Cfg.MONGO_PORT)

	credential := options.Credential{
		AuthSource: "admin",
		Username:   Cfg.MONGO_USER,
		Password:   Cfg.MONGO_PASS,
	}

	var client *mongo.Client
	var err error

	if Cfg.MODE == "PRODUCTION" {

		client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(strConnection).SetAuth(credential))
	} else {

		client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(strConnection))
	}

	if err != nil {
		log.Fatalln(err)
	}

	DB = client
	log.Println("MongoDB connected")
}
