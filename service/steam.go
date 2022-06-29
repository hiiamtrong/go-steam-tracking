package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hiiamtrong/golang-steam-tracking/model"
)

type Response struct {
	Applist struct {
		Apps []model.Game
	}
}

type SteamService struct {
}

func NewSteamService() *SteamService {
	return &SteamService{}
}

func (s *SteamService) GetAllGame() ([]model.Game, error) {
	data, err := http.Get("https://api.steampowered.com/ISteamApps/GetAppList/v0002/")

	if err != nil {
		return nil, err
	}

	if data.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: " + data.Status)
	}

	defer data.Body.Close()

	body, err := ioutil.ReadAll(data.Body)

	if err != nil {
		return nil, err
	}

	var res Response

	newError := json.Unmarshal(body, &res)

	if newError != nil {
		return nil, newError
	}

	return res.Applist.Apps, nil

}

func (s *SteamService) GetGameById(id string) (*model.GameDetail, error) {

	data, err := http.Get(fmt.Sprintf("https://store.steampowered.com/api/appdetails?appids=%s", id))

	if err != nil {
		return nil, err
	}

	if data.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: " + data.Status)
	}

	defer data.Body.Close()

	body, err := ioutil.ReadAll(data.Body)

	if err != nil {
		return nil, err
	}

	var res map[string]interface{}

	jsonError := json.Unmarshal(body, &res)

	if jsonError != nil {
		return nil, jsonError
	}

	if res[id] == nil {
		return nil, fmt.Errorf("game not found")
	}

	gameDetailBytes, err := json.Marshal(res[id].(map[string]interface{})["data"])

	if err != nil {
		return nil, err
	}

	var gameDetail model.GameDetail

	jsonError = json.Unmarshal(gameDetailBytes, &gameDetail)

	if jsonError != nil {
		return nil, jsonError
	}

	return &gameDetail, nil
}

func (s *SteamService) GetGamesPrice(games []model.Game) (map[string]model.PriceOverview, error) {
	var appIds = ""
	for _, game := range games {
		appIds += fmt.Sprintf("%d,", game.AppId)
	}
	res, err := http.Get(fmt.Sprintf("https://store.steampowered.com/api/appdetails?appids=%s&filters=price_overview", appIds))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var resMap map[string]interface{}
	jsonError := json.Unmarshal(body, &resMap)

	if jsonError != nil {
		return nil, jsonError
	}

	result := make(map[string]model.PriceOverview)

	for _, game := range games {
		gameById := resMap[fmt.Sprint(game.AppId)].(map[string]interface{})
		isSuccess := gameById["success"]
		_, ok := gameById["data"].(map[string]interface{})
		if isSuccess == true && ok {
			var priceOverview model.PriceOverview
			priceOverview.FromMap(gameById["data"].(map[string]interface{})["price_overview"].(map[string]interface{}))
			result[fmt.Sprint(game.AppId)] = priceOverview
		}
	}

	return result, nil

}
