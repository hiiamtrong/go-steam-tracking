package repository

import (
	"context"

	"github.com/hiiamtrong/golang-steam-tracking/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GameRepository struct {
	Client *mongo.Client
}

func NewGameRepository(client *mongo.Client) *GameRepository {
	return &GameRepository{
		Client: client,
	}
}

func (r *GameRepository) GetAllGame() ([]model.Game, error) {

	gameCol := r.Client.Database("steam_tracking").Collection("games")

	cur, err := gameCol.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

	var games []model.Game

	for cur.Next(context.TODO()) {
		var game model.Game
		err := cur.Decode(&game)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}

	return games, nil
}
