package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PriceOverview struct {
	Initial          float64 `json:"initial" bson:"initial"`
	Final            float64 `json:"final" bson:"final"`
	Discount         float64 `json:"discount_percent" bson:"discount_percent"`
	InitialFormatted string  `json:"initial_formatted" bson:"initial_formatted"`
	FinalFormatted   string  `json:"final_formatted" bson:"final_formatted"`
	Currency         string  `json:"currency" bson:"currency"`
}

func (p *PriceOverview) FromMap(m map[string]interface{}) {
	p.Initial = m["initial"].(float64)
	p.Final = m["final"].(float64)
	p.Discount = m["discount_percent"].(float64)
	p.InitialFormatted = m["initial_formatted"].(string)
	p.FinalFormatted = m["final_formatted"].(string)
	p.Currency = m["currency"].(string)
}

func (p *PriceOverview) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"initial":           p.Initial,
		"final":             p.Final,
		"discount":          p.Discount,
		"initial_formatted": p.InitialFormatted,
		"final_formatted":   p.FinalFormatted,
		"currency":          p.Currency,
	}
}

type Platforms struct {
	Windows bool `json:"windows" bson:"windows"`
	Mac     bool `json:"mac" bson:"mac"`
	Linux   bool `json:"linux" bson:"linux"`
}

type Category struct {
	Id          int    `json:"id" bson:"id"`
	Description string `json:"description" bson:"description"`
}

type ReleaseDate struct {
	ComingSoon bool   `json:"coming_soon" bson:"coming_soon"`
	Date       string `json:"date" bson:"date"`
}

type GameDetail struct {
	Id               primitive.ObjectID `json:"_id" bson:"_id"`
	SteamAppId       int                `json:"steam_appid" bson:"steam_appid"`
	Name             string             `json:"name" bson:"name"`
	ShortDescription string             `json:"short_description" bson:"short_description"`
	HeaderImage      string             `json:"header_image" bson:"header_image"`
	Platforms        Platforms          `json:"platforms" bson:"platforms"`
	Categories       []Category         `json:"categories" bson:"categories"`
	PriceOverview    *PriceOverview     `json:"price_overview" bson:"price_overview"`
	ReleaseDate      ReleaseDate        `json:"release_date" bson:"release_date"`

	IsFree bool `json:"is_free" bson:"is_free"`
}

func (g *GameDetail) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"steam_appid":       g.SteamAppId,
		"name":              g.Name,
		"short_description": g.ShortDescription,
		"header_image":      g.HeaderImage,
		"platforms":         g.Platforms,
		"categories":        g.Categories,
		"price_overview":    g.PriceOverview,
		"release_date":      g.ReleaseDate,
		"is_free":           g.IsFree,
	}
}

func (g *GameDetail) ToBson() bson.D {
	return bson.D{
		{
			Key:   "steam_appid",
			Value: g.SteamAppId,
		},
		{
			Key:   "name",
			Value: g.Name,
		},
		{
			Key:   "short_description",
			Value: g.ShortDescription,
		},
		{
			Key:   "header_image",
			Value: g.HeaderImage,
		},
		{
			Key:   "platforms",
			Value: g.Platforms,
		},
		{
			Key:   "categories",
			Value: g.Categories,
		},
		{
			Key:   "price_overview",
			Value: g.PriceOverview,
		},
		{
			Key:   "release_date",
			Value: g.ReleaseDate,
		},
		{
			Key:   "is_free",
			Value: g.IsFree,
		},
	}
}
