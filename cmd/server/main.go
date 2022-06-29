package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hiiamtrong/golang-steam-tracking/config"
	"github.com/hiiamtrong/golang-steam-tracking/handler"
	"github.com/hiiamtrong/golang-steam-tracking/repository"
)

func main() {
	config.Setup()
	repository.Setup()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	r.Use(cors.New(config))
	handler.SetupRouter(r)
	log.Println("Server running on port 3001")
	log.Fatalln(r.Run(":3001"))
}
