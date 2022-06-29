package repository

import (
	"context"

	"github.com/hiiamtrong/golang-steam-tracking/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GameDetailRepository struct {
	Client *mongo.Client
}

func NewGameDetailRepository(client *mongo.Client) *GameDetailRepository {
	return &GameDetailRepository{
		Client: client,
	}
}

func (r *GameDetailRepository) GetAllGameDetail(opts *options.FindOptions) ([]model.GameDetailLightW, error) {

	gameDetailCol := r.Client.Database("steam_tracking").Collection("game_details")

	cur, err := gameDetailCol.Find(context.TODO(), bson.M{
		"steam_appid": bson.M{
			"$ne": 0,
		},
	}, opts)

	if err != nil {
		return nil, err
	}

	var gameDetails []model.GameDetailLightW

	for cur.Next(context.TODO()) {
		var gameDetail model.GameDetailLightW
		err := cur.Decode(&gameDetail)
		if err != nil {
			return nil, err
		}
		gameDetails = append(gameDetails, gameDetail)
	}

	return gameDetails, nil
}

func (r *GameDetailRepository) InsertOneGameDetail(gameDetail model.GameDetail) error {

	gameDetailCol := r.Client.Database("steam_tracking").Collection("game_details")

	_, err := gameDetailCol.InsertOne(context.TODO(), gameDetail.ToBson())

	if err != nil {
		return err
	}

	return nil
}

func (r *GameDetailRepository) InsertManyGameDetail(gameDetails []model.GameDetail) error {

	gameDetailCol := r.Client.Database("steam_tracking").Collection("game_details")

	var docs []interface{}

	for _, gameDetail := range gameDetails {
		docs = append(docs, gameDetail.ToBson())
	}

	_, err := gameDetailCol.InsertMany(context.TODO(), docs)

	if err != nil {
		return err
	}

	return nil
}

func (r *GameDetailRepository) SearchGameDetail(query string, opts *options.FindOptions) ([]model.GameDetailLightW, error) {

	gameDetailCol := r.Client.Database("steam_tracking").Collection("game_details")

	var filter bson.M

	if query != "" {
		filter = bson.M{
			"$text": bson.M{
				"$search": query,
			},
		}
	}

	cur, err := gameDetailCol.Find(context.TODO(), filter, opts)

	if err != nil {
		return nil, err
	}

	var gameDetails []model.GameDetailLightW

	for cur.Next(context.TODO()) {
		var gameDetail model.GameDetailLightW
		err := cur.Decode(&gameDetail)
		if err != nil {
			return nil, err
		}
		gameDetails = append(gameDetails, gameDetail)
	}

	return gameDetails, nil
}

func (r *GameDetailRepository) GetGameDetailById(id int) (model.GameDetail, error) {

	gameDetailCol := r.Client.Database("steam_tracking").Collection("game_details")

	var gameDetail model.GameDetail

	err := gameDetailCol.FindOne(context.TODO(), bson.M{
		"steam_appid": id,
	}).Decode(&gameDetail)

	if err != nil {
		return model.GameDetail{}, err
	}

	return gameDetail, nil
}

func (r *GameDetailRepository) UpdatePriceOverview(id int, priceOverview *model.PriceOverview) error {

	gameDetailCol := r.Client.Database("steam_tracking").Collection("game_details")

	_, err := gameDetailCol.UpdateOne(context.TODO(), bson.M{
		"steam_appid": id,
	}, bson.M{
		"$set": bson.M{
			"price_overview": priceOverview,
		},
	})

	if err != nil {
		return err
	}

	return nil
}
