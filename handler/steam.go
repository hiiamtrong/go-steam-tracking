package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiiamtrong/golang-steam-tracking/service"
)

type SteamHandler struct {
	steamService *service.SteamService
}

func NewSteamHandler(steamService *service.SteamService) *SteamHandler {
	return &SteamHandler{
		steamService: steamService,
	}
}

func (s *SteamHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		games, err := s.steamService.GetAllGame()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error()})
			return

		}
		c.JSON(http.StatusOK, games)
	}
}

func (s *SteamHandler) GetGameDetailById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		game, err := s.steamService.GetGameById(id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error()})
			return

		}
		c.JSON(http.StatusOK, game.ToMap())
	}
}
