package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hiiamtrong/golang-steam-tracking/model"
)

type ResponseV2 struct {
	Applist struct {
		Apps []model.Game
	}
}

type ResponseV1 struct {
	Applist struct {
		Apps struct {
			App []model.Game
		}
	}
}

type SteamService struct {
}

func NewSteamService() *SteamService {
	return &SteamService{}
}

func (s *SteamService) GetAllGame(version string) ([]model.Game, error) {
	data, err := http.Get("https://api.steampowered.com/ISteamApps/GetAppList/" + version)

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

	var resv1 ResponseV1
	var resv2 ResponseV2

	if version == "v2" {

		newError := json.Unmarshal(body, &resv2)
		if newError != nil {
			return nil, newError
		}

		return resv2.Applist.Apps, nil
	} else {
		newError := json.Unmarshal(body, &resv1)
		if newError != nil {
			return nil, newError
		}

		return resv1.Applist.Apps.App, nil
	}

}

func (s *SteamService) GetGameById(id string) (*model.GameDetail, error) {

	data, err := http.Get(fmt.Sprintf("https://store.steampowered.com/api/appdetails?appids=%s&cc=vnd", id))

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

func (s *SteamService) GetGamesPrice(games []model.GameDetailLightW) (map[string]*model.PriceOverview, error) {
	var appIds = ""
	for _, game := range games {
		appIds += fmt.Sprintf("%d,", game.SteamAppId)
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

	if resMap == nil {
		return nil, fmt.Errorf("game not found")
	}

	result := make(map[string]*model.PriceOverview)

	for _, game := range games {
		gameById := resMap[fmt.Sprint(game.SteamAppId)].(map[string]interface{})
		isSuccess := gameById["success"]
		_, ok := gameById["data"].(map[string]interface{})
		if isSuccess == true && ok {
			var priceOverview model.PriceOverview
			priceOverview.FromMap(gameById["data"].(map[string]interface{})["price_overview"].(map[string]interface{}))
			result[fmt.Sprint(game.SteamAppId)] = &priceOverview
		}
	}

	return result, nil

}
