package repository

import "github.com/hiiamtrong/golang-steam-tracking/config"

var GameRepositoryInstance *GameRepository
var GameDetailRepositoryInstance *GameDetailRepository

func Setup() {
	GameRepositoryInstance = NewGameRepository(config.DB)
	GameDetailRepositoryInstance = NewGameDetailRepository(config.DB)
}
