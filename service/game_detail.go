package service

import (
	"github.com/hiiamtrong/golang-steam-tracking/dto"
	"github.com/hiiamtrong/golang-steam-tracking/model"
	"github.com/hiiamtrong/golang-steam-tracking/repository"
	"github.com/samber/lo"
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

func (s *GameDetailService) GetAllGameDetail(req dto.GetAllGameDetailRequest) ([]model.GameDetail, error) {
	opts := options.Find()

	opts.SetLimit(int64(req.Limit))
	opts.SetSkip(int64(req.Page * req.Limit))

	gameDetails, err := s.GameDetailRepository.GetAllGameDetail(opts)

	if err != nil {
		return nil, err
	}

	filter := lo.Filter(gameDetails, func(g model.GameDetail, _ int) bool {
		return g.SteamAppId != 0
	})

	return filter, nil
}

func (s *GameDetailService) SearchGameDetail(req dto.SearchGameDetailRequest) ([]model.GameDetail, error) {
	opts := options.Find()

	opts.SetLimit(int64(req.Limit))
	opts.SetSkip(int64(req.Page * req.Limit))

	gameDetails, err := s.GameDetailRepository.SearchGameDetail(req.Query, opts)

	if err != nil {
		return nil, err
	}

	filter := lo.Filter(gameDetails, func(g model.GameDetail, _ int) bool {
		return g.SteamAppId != 0
	})

	return filter, nil
}
