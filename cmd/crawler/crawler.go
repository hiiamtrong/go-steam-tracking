package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hiiamtrong/golang-steam-tracking/config"
	"github.com/hiiamtrong/golang-steam-tracking/repository"
	"github.com/hiiamtrong/golang-steam-tracking/service"
	"github.com/hiiamtrong/golang-steam-tracking/util"
	"github.com/samber/lo"

	"github.com/jasonlvhit/gocron"
)

func main() {

	config.Setup()
	util.Setup()
	repository.Setup()

	s := gocron.NewScheduler()

	s.Every(1).Day().At("22:35").Do(insertAllGameTask)

	s.Every(1).Day().At("22:35").Do(insertAllGameDetailTask)
	insertAllGameTask()
	insertAllGameDetailTask()

	<-s.Start()

}

func insertAllGameDetailTask() {
	log.Println("Start insert all games detail")

	games, err := repository.GameRepositoryInstance.GetAllGame()

	if err != nil {
		panic(err)
	}

	gameDetails, er := repository.GameDetailRepositoryInstance.GetAllGameDetail(nil)

	if er != nil {
		panic(er)
	}

	var gameDetailAppIds []int
	var gameAppIds []int

	for _, game := range games {
		gameAppIds = append(gameAppIds, game.AppId)
	}

	for _, gameDetail := range gameDetails {
		gameDetailAppIds = append(gameDetailAppIds, gameDetail.SteamAppId)

	}

	fmt.Println(len(gameDetailAppIds), len(gameAppIds))

	left, right := lo.Difference(gameAppIds, gameDetailAppIds)

	needToInsert := append(left, right...)

	fmt.Println(len(left), len(right), len(needToInsert))

	steamService := service.NewSteamService()

	for _, appId := range needToInsert {
		log.Println("Crawling game detail for appId: ", appId)
		gameDetail, err := steamService.GetGameById(fmt.Sprint(appId))

		if err != nil {

			log.Println("Error when get game detail: ", err)
			if err.Error() == "Error: 429 Too Many Requests" {

				fmt.Println("Sleep for 1 minute")
				time.Sleep(time.Second * 60)

				continue
			}

			continue
		}

		if gameDetail.SteamAppId == 0 {
			continue
		}

		err = repository.GameDetailRepositoryInstance.InsertOneGameDetail(*gameDetail)

		if err != nil {
			log.Println("Error when insert game detail: ", err)
			continue
		}

		log.Println("Inserted game detail for appId: ", appId)

		time.Sleep(500 * time.Millisecond)

	}

}

func insertAllGameTask() {
	log.Println("Start insert all games")
	gameCol := config.DB.Database("steam_tracking").Collection("games")

	log.Println("Droping games collection...")
	log.Println("==============================")
	var err = gameCol.Drop(context.TODO())
	if err != nil {
		panic(err)
	}
	log.Println("Dropped games collection")
	log.Println("==============================")

	log.Println("Inserting all games...")
	log.Println("==============================")

	steamService := service.NewSteamService()

	games, err := steamService.GetAllGame()

	if err != nil {
		panic(err)
	}

	var docs []interface{}

	for _, game := range games {
		docs = append(docs, game.ToBson())
	}

	_, err = gameCol.InsertMany(context.TODO(), docs)

	if err != nil {
		panic(err)
	}

	log.Println("Inserted all games")
}
