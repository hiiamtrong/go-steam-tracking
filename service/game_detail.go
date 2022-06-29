package service

import (
	"github.com/hiiamtrong/golang-steam-tracking/dto"
	"github.com/hiiamtrong/golang-steam-tracking/model"
	"github.com/hiiamtrong/golang-steam-tracking/repository"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GameDetailService struct {
	GameDetailRepository *repository.GameDetailRepository
}

func NewGameDetailService(gameDetailRepository *repository.GameDetailRepository) *GameDetailService {
	return &GameDetailService{
		GameDetailRepository: gameDetailRepository,
	}
}

func (s *GameDetailService) GetAllGameDetail(req dto.GetAllGameDetailRequest) ([]model.GameDetailLightW, error) {
	opts := options.Find()

	opts.SetLimit(int64(req.Limit))
	opts.SetSkip(int64(req.Page * req.Limit))

	gameDetails, err := s.GameDetailRepository.GetAllGameDetail(opts)

	if err != nil {
		return nil, err
	}

	filter := lo.Filter(gameDetails, func(g model.GameDetailLightW, _ int) bool {
		return g.SteamAppId != 0 && g.Type == "game"
	})

	return filter, nil
}

func (s *GameDetailService) SearchGameDetail(req dto.SearchGameDetailRequest) ([]model.GameDetailLightW, error) {
	opts := options.Find()

	opts.SetLimit(int64(req.Limit))
	opts.SetSkip(int64(req.Page * req.Limit))

	switch req.SortBy {
	case "name":
		opts.SetSort(bson.D{{Key: "name", Value: req.SortOrder}, {Key: "price_overview.discount_percent", Value: -1}})
	case "price":
		opts.SetSort(bson.D{{Key: "price_overview.final", Value: req.SortOrder}})
	case "release_date":
		opts.SetSort(bson.D{{Key: "release_date.date", Value: req.SortOrder}, {Key: "price_overview.discount_percent", Value: -1}})
	case "discount":
		opts.SetSort(bson.D{{Key: "price_overview.discount_percent", Value: req.SortOrder}, {
			Key: "price_overview.final", Value: -req.SortOrder}})
	case "recommend":
		opts.SetSort(bson.D{{Key: "recommendations.total", Value: req.SortOrder}, {Key: "price_overview.discount_percent", Value: -1}})
	case "metacritic":
		opts.SetSort(bson.D{{Key: "metacritic.score", Value: req.SortOrder}, {Key: "price_overview.discount_percent", Value: -1}})
	}

	gameDetails, err := s.GameDetailRepository.SearchGameDetail(req.Query, opts)

	if err != nil {
		return nil, err
	}

	filter := lo.Filter(gameDetails, func(g model.GameDetailLightW, _ int) bool {
		return g.SteamAppId != 0 && g.Type == "game"
	})

	return filter, nil
}

func (s *GameDetailService) GetGameDetailById(req dto.GetGameDetailByIdRequest) (model.GameDetail, error) {
	return s.GameDetailRepository.GetGameDetailById(req.Id)
}
