package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hiiamtrong/golang-steam-tracking/repository"
	"github.com/hiiamtrong/golang-steam-tracking/service"
)

func SetupRouter(app *gin.Engine) {

	gameDetailHandler := NewGameDetailHandler(service.NewGameDetailService(repository.GameDetailRepositoryInstance))
	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Pong",
		})
	})
	app.GET("/game_detail", gameDetailHandler.GetAllGameDetail())
	app.GET("/game_detail/:id", gameDetailHandler.GetGameDetailById())
	app.GET("/game_detail/search", gameDetailHandler.SearchGameDetail())
}
