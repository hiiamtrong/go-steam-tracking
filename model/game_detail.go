package model

import (
	"log"

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

type Metacritic struct {
	Score int    `json:"score" bson:"score"`
	Url   string `json:"url" bson:"url"`
}

type Genre struct {
	Id          string `json:"id" bson:"id"`
	Description string `json:"description" bson:"description"`
}

type Screenshot struct {
	Id            int    `json:"id" bson:"id"`
	PathThumbnail string `json:"path_thumbnail" bson:"path_thumbnail"`
	PathFull      string `json:"path_full" bson:"path_full"`
}

type Movie struct {
	Id        int    `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	Thumbnail string `json:"thumbnail" bson:"thumbnail"`
	Webm      struct {
		P480 string `json:"480" bson:"480"`
		Max  string `json:"max" bson:"max"`
	}

	Mp4 struct {
		P480 string `json:"480" bson:"480"`
		Max  string `json:"max" bson:"max"`
	}

	Highlight bool `json:"highlight" bson:"highlight"`
}

type Recommendation struct {
	Total int `json:"total" bson:"total"`
}
type SupportInfo struct {
	Url   string `json:"url" bson:"url"`
	Email string `json:"email" bson:"email"`
}

type GameDetail struct {
	Id                 primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	SteamAppId         int                `json:"steam_appid" bson:"steam_appid" `
	Name               string             `json:"name" bson:"name"`
	Type               string             `json:"type" bson:"type"`
	SupportedLanguages string             `json:"supported_languages" bson:"supported_languages"`
	AboutTheGame       string             `json:"about_the_game" bson:"about_the_game"`

	RequiredAge         interface{} `json:"required_age" bson:"required_age"`
	Dlc                 []int       `json:"dlc" bson:"dlc"`
	DetailedDescription string      `json:"detailed_description" bson:"detailed_description"`
	PcRequirements      interface{} `json:"pc_requirements" bson:"pc_requirements"`
	MacRequirements     interface{} `json:"mac_requirements" bson:"mac_requirements"`
	LinuxRequirements   interface{} `json:"linux_requirements" bson:"linux_requirements"`
	LegalNotice         string      `json:"legal_notice" bson:"legal_notice"`
	Developers          []string    `json:"developers" bson:"developers"`
	Publishers          []string    `json:"publishers" bson:"publishers"`
	Packages            []int       `json:"packages" bson:"packages"`

	Metacritic      *Metacritic     `json:"metacritic" bson:"metacritic"`
	Genres          []Genre         `json:"genres" bson:"genres"`
	Screenshots     []Screenshot    `json:"screenshots" bson:"screenshots"`
	Movies          []Movie         `json:"movies" bson:"movies"`
	Recommendations *Recommendation `json:"recommendations" bson:"recommendations"`
	SupportInfo     *SupportInfo    `json:"support_info" bson:"support_info"`
	Background      string          `json:"background" bson:"background"`
	BackgroundRaw   string          `json:"background_raw" bson:"background_raw"`

	ControllerSupport string         `json:"controller_support" bson:"controller_support"`
	ShortDescription  string         `json:"short_description" bson:"short_description"`
	HeaderImage       string         `json:"header_image" bson:"header_image"`
	Platforms         *Platforms     `json:"platforms" bson:"platforms"`
	Categories        []Category     `json:"categories" bson:"categories"`
	PriceOverview     *PriceOverview `json:"price_overview" bson:"price_overview"`
	ReleaseDate       *ReleaseDate   `json:"release_date" bson:"release_date"`

	IsFree bool `json:"is_free" bson:"is_free"`
}

func (g *GameDetail) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"_id":                  g.Id,
		"steam_appid":          g.SteamAppId,
		"name":                 g.Name,
		"type":                 g.Type,
		"required_age":         g.RequiredAge,
		"dlc":                  g.Dlc,
		"detailed_description": g.DetailedDescription,
		"pc_requirements":      g.PcRequirements,
		"mac_requirements":     g.MacRequirements,
		"linux_requirements":   g.LinuxRequirements,
		"legal_notice":         g.LegalNotice,
		"developers":           g.Developers,
		"publishers":           g.Publishers,
		"packages":             g.Packages,
		"metacritic":           g.Metacritic,
		"genres":               g.Genres,
		"screenshots":          g.Screenshots,
		"movies":               g.Movies,
		"recommendations":      g.Recommendations,
		"support_info":         g.SupportInfo,
		"background":           g.Background,
		"background_raw":       g.BackgroundRaw,
		"controller_support":   g.ControllerSupport,
		"short_description":    g.ShortDescription,
		"header_image":         g.HeaderImage,
		"platforms":            g.Platforms,
		"categories":           g.Categories,
		"price_overview":       g.PriceOverview,
		"release_date":         g.ReleaseDate,
		"is_free":              g.IsFree,
	}
}

func (g *GameDetail) ToBson() *bson.D {
	data, err := bson.Marshal(g)
	if err != nil {
		log.Println(err)
	}
	var result bson.D
	err = bson.Unmarshal(data, &result)
	if err != nil {
		log.Println(err)
	}

	return &result
}

type GameDetailLightW struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	SteamAppId  int                `json:"steam_appid" bson:"steam_appid" `
	Name        string             `json:"name" bson:"name"`
	Type        string             `json:"type" bson:"type"`
	HeaderImage string             `json:"header_image" bson:"header_image"`
	IsFree      bool               `json:"is_free" bson:"is_free"`
	Platforms   *Platforms         `json:"platforms" bson:"platforms"`

	PriceOverview *PriceOverview `json:"price_overview" bson:"price_overview"`
	ReleaseDate   *ReleaseDate   `json:"release_date" bson:"release_date"`
}
