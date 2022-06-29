package repository

import (
	"context"

	"github.com/hiiamtrong/golang-steam-tracking/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlackListRepository struct {
	Client *mongo.Client
}

func NewBlackListRepository(client *mongo.Client) *BlackListRepository {
	return &BlackListRepository{
		Client: client,
	}
}

func (r *BlackListRepository) GetAllBlackList(opts *options.FindOptions) ([]model.BlackList, error) {

	blackListCol := r.Client.Database("steam_tracking").Collection("black_list")

	cur, err := blackListCol.Find(context.TODO(), bson.M{}, opts)

	if err != nil {
		return nil, err
	}

	var blackLists []model.BlackList

	for cur.Next(context.TODO()) {
		var blackList model.BlackList
		err := cur.Decode(&blackList)
		if err != nil {
			return nil, err
		}
		blackLists = append(blackLists, blackList)
	}

	return blackLists, nil
}

func (r *BlackListRepository) InsertOneBlackList(blackList model.BlackList) error {

	blackListCol := r.Client.Database("steam_tracking").Collection("black_list")

	_, err := blackListCol.InsertOne(context.TODO(), blackList.ToBson())

	if err != nil {
		return err
	}

	return nil
}
