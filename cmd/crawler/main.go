package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"time"

	"github.com/hiiamtrong/golang-steam-tracking/config"
	"github.com/hiiamtrong/golang-steam-tracking/model"
	"github.com/hiiamtrong/golang-steam-tracking/repository"
	"github.com/hiiamtrong/golang-steam-tracking/service"
	"github.com/hiiamtrong/golang-steam-tracking/util"
	"github.com/robfig/cron/v3"
	"github.com/samber/lo"
)

func main() {

	config.Setup()
	util.Setup()
	repository.Setup()

	sig := make(chan os.Signal)

	c := cron.New()

	c.AddFunc("10 00 * * *", insertAllGameTask)
	c.AddFunc("15 00 * * *", insertAllGameDetailTask)
	c.AddFunc("30 * * * *", crawlPriceHistory)

	go c.Start()
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
	fmt.Println("Cron job stopped")

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

	blackLists, err := repository.BlackListRepositoryInstance.GetAllBlackList(nil)
	if err != nil {
		panic(err)
	}

	var total []int
	var existIds []int

	for _, game := range games {
		total = append(total, game.AppId)
	}
	for _, gameDetail := range gameDetails {
		existIds = append(existIds, gameDetail.SteamAppId)
	}
	for _, blackList := range blackLists {
		existIds = append(existIds, blackList.SteamAppId)
	}

	removeDuplicateIds(existIds)

	fmt.Println("Exists len: ", len(existIds))
	fmt.Println("Total game length: ", len(total))
	left, right := lo.Difference(total, existIds)

	fmt.Println("Left: ", len(left))
	fmt.Println("Right: ", len(right))

	needToInsert := lo.Shuffle(left)

	fmt.Println(len(needToInsert))

	steamService := service.NewSteamService()

	for _, appId := range needToInsert {
		log.Println("Crawling game detail for appId: ", appId)

		gameDetail, err := steamService.GetGameById(fmt.Sprint(appId))

		if err != nil {

			log.Println("Error when get game detail: ", err)
			if err.Error() == "Error: 429 Too Many Requests" {

				needToInsert = append(needToInsert, appId)
				fmt.Println("Sleep for 1 minute")
				time.Sleep(time.Second * 60)

				continue
			}

			if err.Error() == "unexpected end of JSON input" {
				repository.BlackListRepositoryInstance.InsertOneBlackList(
					model.BlackList{
						SteamAppId: appId},
				)
			}

			continue
		}

		if gameDetail.SteamAppId == 0 {
			repository.BlackListRepositoryInstance.InsertOneBlackList(
				model.BlackList{
					SteamAppId: appId},
			)
			continue
		}

		err = repository.GameDetailRepositoryInstance.InsertOneGameDetail(*gameDetail)

		if err != nil {
			log.Println("Error when insert game detail: ", err)
			// regexp to check if error is because of duplicate key
			regexp := regexp.MustCompile(`duplicate key error collection`)

			if regexp.MatchString(err.Error()) {
				log.Println("Game detail already exist")
				repository.BlackListRepositoryInstance.InsertOneBlackList(
					model.BlackList{
						SteamAppId: appId,
					},
				)
			}

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

	gamesV1, err := steamService.GetAllGame("v1")
	if err != nil {
		panic(err)
	}

	fmt.Println(len(gamesV1))
	gamesV2, err := steamService.GetAllGame("v2")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(gamesV2))

	maps := make(map[int]model.Game)

	var total []model.Game

	for _, game := range gamesV1 {
		maps[game.AppId] = game
		total = append(total, game)
	}

	for _, game := range gamesV2 {
		if _, ok := maps[game.AppId]; !ok {
			fmt.Println("AppId: ", game.AppId)
			maps[game.AppId] = game
			total = append(total, game)
		}
	}

	if err != nil {
		panic(err)
	}

	var docs []interface{}

	for _, game := range total {
		docs = append(docs, game.ToBson())
	}

	_, err = gameCol.InsertMany(context.TODO(), docs)

	if err != nil {
		panic(err)
	}

	log.Println("Inserted all games")
}

func removeDuplicateIds(ids []int) []int {
	log.Println("Start remove duplicate ids")

	mapDup := make(map[int]int)

	for _, id := range ids {
		if _, ok := mapDup[id]; !ok {
			mapDup[id] = 0
		}
		mapDup[id]++
	}

	var newIds []int

	for id, count := range mapDup {
		if count == 1 {
			newIds = append(newIds, id)
		}
	}

	return newIds

}

func crawlPriceHistory() {
	log.Println("Start crawl price history")
	gameDetails, err := repository.GameDetailRepositoryInstance.GetAllGameDetail(nil)

	steamService := service.NewSteamService()
	if err != nil {
		panic(err)
	}

	chunkGameDetails := lo.Chunk(gameDetails, 1000)

	rs := make(map[string]*model.PriceOverview)
	for _, chunk := range chunkGameDetails {

		chunkResult, err := steamService.GetGamesPrice(chunk)
		if err != nil {
			log.Println("History Error:", err.Error())
			continue
		}
		rs = mergePriceHistory(rs, chunkResult)
	}

	var priceHistoryDocs []model.PriceHistory
	for _, game := range gameDetails {
		if _, ok := rs[fmt.Sprint(game.SteamAppId)]; !ok {
			continue
		}
		if comparePriceOverview(game.PriceOverview, rs[fmt.Sprint(game.SteamAppId)]) {
			continue
		}

		priceHistoryDocs = append(priceHistoryDocs, model.PriceHistory{
			SteamAppId:    game.SteamAppId,
			PriceOverview: rs[fmt.Sprint(game.SteamAppId)],
			Date:          time.Now(),
		})

	}

	if len(priceHistoryDocs) <= 0 {
		log.Println("No price history to update")
		return
	}

	fmt.Println("Total price apps updating: ", len(priceHistoryDocs))
	err = repository.PriceHistoryRepositoryInstance.InsertManyPriceHistory(priceHistoryDocs)

	if err != nil {
		log.Println("Error when insert price history: ", err)
	}

	for _, game := range priceHistoryDocs {
		err = repository.GameDetailRepositoryInstance.UpdatePriceOverview(game.SteamAppId, game.PriceOverview)
		if err != nil {
			log.Println("Error when update price overview: ", err)
		}
	}

	fmt.Println("Crawled price history")

}

func mergePriceHistory(rs map[string]*model.PriceOverview, result map[string]*model.PriceOverview) map[string]*model.PriceOverview {
	for key, value := range result {
		if _, ok := rs[key]; !ok {
			rs[key] = value
		} else {
			log.Println("Duplicate key: ", key)
		}
	}
	return rs
}

func comparePriceOverview(origin *model.PriceOverview, new *model.PriceOverview) bool {
	if origin == nil && new == nil {
		return true
	}
	if origin == nil || new == nil {
		return false
	}

	if origin.Initial != new.Initial {
		return false
	}
	if origin.Final != new.Final {
		return false
	}

	if origin.Currency != new.Currency {
		return false
	}

	return true
}
