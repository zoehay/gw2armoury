package gw2api

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/zoehay/gw2armoury/backend/internal/clients"
	apimodels "github.com/zoehay/gw2armoury/backend/internal/gw2_api/api_models"
)

func GetAllCharacters(apiKey string) ([]apimodels.APICharacter, error) {
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

	var result []apimodels.APICharacter

	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("provider json.Unmarshal error: %s", err)
	}

	return result, nil
}

//	var mockBagItems = []ApiBagItem{
//		{
//			"id":      69432,
//			"count":   62,
//			"binding": "Account",
//		},
//		{
//			"id":      92209,
//			"count":   1,
//			"binding": "Account",
//		},
//		{
//			"id":      92209,
//			"count":   1,
//			"binding": "Account",
//		},
//	}

// var bindingType = "Character"
// var boundTo = "Character Name"

// var bagItems = []*apimodels.ApiBagItem{
// 	{
// 		Id:      32130,
// 		Count:   1,
// 		Binding: &bindingType,
// 		BoundTo: &boundTo,
// 	},
// 	{
// 		Id:    32130,
// 		Count: 1,
// 	},
// 	{
// 		Id:    32130,
// 		Count: 1,
// 	},
// 	{
// 		Id:      32130,
// 		Count:   1,
// 		Binding: &bindingType,
// 		BoundTo: &boundTo,
// 	},
// }

// var apiBags = []apimodels.ApiBag{
// 	{
// 		Id:        8932,
// 		Size:      20,
// 		Inventory: bagItems,
// 	},
// 	{
// 		Id:        8936,
// 		Size:      15,
// 		Inventory: bagItems,
// 	},
// }

// var character = []apimodels.APICharacter{
// 	{
// 		Name: "Laura Lesdottir",
// 		Bags: &apiBags,
// 	},
// 	{
// 		Name: "Other Name",
// 		Bags: &apiBags,
// 	},
// }
