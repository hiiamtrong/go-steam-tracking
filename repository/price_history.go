package repository

import (
	"context"

	"github.com/hiiamtrong/golang-steam-tracking/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PriceHistoryRepository struct {
	Client *mongo.Client
}

func NewPriceHistoryRepository(client *mongo.Client) *PriceHistoryRepository {
	return &PriceHistoryRepository{
		Client: client,
	}
}

func (r *PriceHistoryRepository) GetBestPrice(appId int) (*model.PriceHistory, error) {

	priceHistoryCol := r.Client.Database("steam_tracking").Collection("price_history")

	var priceHistories []model.PriceHistory

	cur, err := priceHistoryCol.Aggregate(
		context.TODO(),
		[]bson.M{
			{
				"$match": bson.M{
					"steam_appid": appId,
				},
			},
			{
				"$group": bson.M{
					"_id": "$steam_appid",
					"min": bson.M{
						"$min": "$price_overview.final",
					},
				},
			},
		},
	)

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var priceHistory model.PriceHistory
		err := cur.Decode(&priceHistory)
		if err != nil {
			return nil, err
		}
		priceHistories = append(priceHistories, priceHistory)
	}

	return &priceHistories[0], nil
}

func (r PriceHistoryRepository) InsertManyPriceHistory(priceHistories []model.PriceHistory) error {

	priceHistoryCol := r.Client.Database("steam_tracking").Collection("price_history")

	var priceHistoriesBson []interface{}

	for _, priceHistory := range priceHistories {
		priceHistoriesBson = append(priceHistoriesBson, priceHistory.ToBson())
	}

	_, err := priceHistoryCol.InsertMany(context.TODO(), priceHistoriesBson)

	if err != nil {
		return err
	}

	return nil
}
