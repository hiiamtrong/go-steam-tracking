package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hiiamtrong/golang-steam-tracking/repository"
	"github.com/hiiamtrong/golang-steam-tracking/service"
)

func SetupRouter(app *gin.Engine) {
	steamHandler := NewSteamHandler(service.NewSteamService())
	app.GET("/", steamHandler.GetAll())
	app.GET("/:id", steamHandler.GetGameDetailById())

	gameDetailHandler := NewGameDetailHandler(service.NewGameDetailService(repository.GameDetailRepositoryInstance))

	app.GET("/game_detail", gameDetailHandler.GetAllGameDetail())
	app.GET("/game_detail/search", gameDetailHandler.SearchGameDetail())
}
