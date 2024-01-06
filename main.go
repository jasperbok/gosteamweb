package gosteamweb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WebApiClient struct {
	ApiKey    string
	ProfileId string
	client    *http.Client
}

func (c WebApiClient) GetOwnedGames() ([]App, error) {
	var apps []App

	url := fmt.Sprintf("https://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json", c.ApiKey, c.ProfileId)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return apps, err
	}

	response, err := c.client.Do(request)
	if err != nil {
		return apps, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return apps, err
	}

	responseStruct := struct {
		Response struct {
			Games []App `json:"games"`
		} `json:"response"`
	}{}

	err = json.Unmarshal(body, &responseStruct)
	if err != nil {
		return apps, err
	}

	return responseStruct.Response.Games, nil
}

func NewWebApiClient(apiKey, profileId string) *WebApiClient {
	return &WebApiClient{ApiKey: apiKey, ProfileId: profileId, client: &http.Client{}}
}
