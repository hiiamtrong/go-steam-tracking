package repository

import "github.com/hiiamtrong/golang-steam-tracking/config"

var (
	GameRepositoryInstance         *GameRepository
	GameDetailRepositoryInstance   *GameDetailRepository
	BlackListRepositoryInstance    *BlackListRepository
	PriceHistoryRepositoryInstance *PriceHistoryRepository
)

func Setup() {
	GameRepositoryInstance = NewGameRepository(config.DB)
	GameDetailRepositoryInstance = NewGameDetailRepository(config.DB)
	BlackListRepositoryInstance = NewBlackListRepository(config.DB)
	PriceHistoryRepositoryInstance = NewPriceHistoryRepository(config.DB)

}
