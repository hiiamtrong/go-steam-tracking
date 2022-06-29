package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Game struct {
	Id    primitive.ObjectID `json:"_id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	AppId int                `json:"appid" bson:"appid"`
}

func (g *Game) ToBson() bson.D {
	return bson.D{
		{Key: "name", Value: g.Name},
		{Key: "appid", Value: g.AppId},
	}
}
