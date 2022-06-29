package service

import (
	"github.com/hiiamtrong/golang-steam-tracking/model"
	"github.com/hiiamtrong/golang-steam-tracking/repository"
)

type GameService struct {
	GameRepository *repository.GameRepository
}

func NewGameService(gameRepository *repository.GameRepository) *GameService {
	return &GameService{
		GameRepository: gameRepository,
	}
}

func (s *GameService) GetAllGame() ([]model.Game, error) {
	return s.GameRepository.GetAllGame()
}
