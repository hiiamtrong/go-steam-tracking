package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PriceHistory struct {
	Id            primitive.ObjectID `json:"id" bson:"_id"`
	SteamAppId    int                `json:"steam_appid" bson:"steam_appid"`
	PriceOverview *PriceOverview     `json:"price_overview" bson:"price_overview"`

	Date time.Time `json:"date" bson:"date"`
}

func (p *PriceHistory) ToBson() bson.D {
	return bson.D{
		{Key: "steam_appid", Value: p.SteamAppId},
		{Key: "price_overview", Value: p.PriceOverview},
		{Key: "date", Value: p.Date},
	}
}
