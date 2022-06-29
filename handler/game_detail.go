package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hiiamtrong/golang-steam-tracking/dto"
	"github.com/hiiamtrong/golang-steam-tracking/service"
)

type GameDetailHandler struct {
	GameDetailService *service.GameDetailService
}

func NewGameDetailHandler(gameDetailService *service.GameDetailService) *GameDetailHandler {
	return &GameDetailHandler{
		GameDetailService: gameDetailService,
	}
}

func (h *GameDetailHandler) GetAllGameDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
		page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))

		gameDetails, err := h.GameDetailService.GetAllGameDetail(
			dto.GetAllGameDetailRequest{
				Limit: limit,
				Page:  page,
			},
		)

		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gameDetails)
	}
}

func (h *GameDetailHandler) SearchGameDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Query("query")
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
		page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
		sortBy := c.DefaultQuery("sort_by", "name")
		sortOrder, _ := strconv.Atoi(c.DefaultQuery("sort_order", "1"))

		gameDetails, err := h.GameDetailService.SearchGameDetail(
			dto.SearchGameDetailRequest{
				Query:     query,
				Limit:     limit,
				Page:      page,
				SortBy:    sortBy,
				SortOrder: sortOrder,
			},
		)

		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gameDetails)
	}
}

func (h *GameDetailHandler) GetGameDetailById() gin.HandlerFunc {
	return func(c *gin.Context) {
		steamAppId, _ := strconv.Atoi(c.Param("id"))

		gameDetail, err := h.GameDetailService.GetGameDetailById(dto.GetGameDetailByIdRequest{Id: steamAppId})

		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})

			return
		}
		c.JSON(200, gameDetail)
	}
}
