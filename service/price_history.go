package service

import (
	"github.com/hiiamtrong/golang-steam-tracking/model"
	"github.com/hiiamtrong/golang-steam-tracking/repository"
)

type PriceHistoryService struct {
	priceHistoryRepository repository.PriceHistoryRepository
}

func NewPriceHistoryService(priceHistoryRepository repository.PriceHistoryRepository) PriceHistoryService {
	return PriceHistoryService{
		priceHistoryRepository: priceHistoryRepository,
	}
}

func (s PriceHistoryService) GetBestPrice(appId int) (*model.PriceHistory, error) {
	return s.priceHistoryRepository.GetBestPrice(appId)
}

func (s PriceHistoryService) InsertManyPriceHistory(priceHistories []model.PriceHistory) error {
	return s.priceHistoryRepository.InsertManyPriceHistory(priceHistories)
}
