package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlackList struct {
	Id         primitive.ObjectID `json:"_id" bson:"_id"`
	SteamAppId int                `json:"steam_appid" bson:"steam_appid"`
}

func (b *BlackList) ToBson() bson.D {
	return bson.D{
		{Key: "steam_appid", Value: b.SteamAppId},
	}
}
