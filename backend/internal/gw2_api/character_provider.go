package gw2api

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/zoehay/gw2armoury/backend/internal/clients"
	apimodels "github.com/zoehay/gw2armoury/backend/internal/gw2_api/api_models"
)

// func GetAllCharacterNames(apiKey string) ([]string, error) {

// }

// var mockAllCharacterNames = []string{}

// func GetCharacterInventory(characterName string, apiKey string) ([]ApiBagItem, error) {
// 	return mockBagItems, nil

// }

// var mockBagItems = []ApiBagItem{
// 	{
// 		"id":      69432,
// 		"count":   62,
// 		"binding": "Account",
// 	},
// 	{
// 		"id":      92209,
// 		"count":   1,
// 		"binding": "Account",
// 	},
// 	{
// 		"id":      92209,
// 		"count":   1,
// 		"binding": "Account",
// 	},
// }

func GetAllCharacters(apiKey string) ([]apimodels.ApiCharacter, error) {
	res, err := clients.GetAllCharacters(apiKey)

	if err != nil {
		return nil, fmt.Errorf("provider get error: %s", err)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("provider io.ReadAll error: %s", err)
	}

	var result []apimodels.ApiCharacter

	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("provider json.Unmarshal error: %s", err)
	}

	return result, nil
}
