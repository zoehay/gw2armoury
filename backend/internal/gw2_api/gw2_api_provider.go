package gw2api

import (
	"fmt"
	"io"
	"strconv"

	"encoding/json"

	"github.com/zoehay/gw2armoury/backend/internal/clients"
	apimodels "github.com/zoehay/gw2armoury/backend/internal/gw2_api/api_models"
)

// getCharacterInventory bags: Bag[]

func GetItemsById(ids string) ([]apimodels.ApiItem, error) {
	res, err := clients.GetItemsById(ids)

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

	var result []apimodels.ApiItem

	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("provider json.Unmarshal error: %s", err)
	}

	return result, nil
}

func GetAllItemIds() ([]string, error) {
	// res, err := clients.GetItemIds()

	// if err != nil {
	// 	return nil, fmt.Errorf("provider get error: %s", err)
	// }

	// defer func() {
	// 	_ = res.Body.Close()
	// }()

	// body, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("provider io.ReadAll error: %s", err)
	// }

	// var result []int

	// if err = json.Unmarshal(body, &result); err != nil {
	// 	return nil, fmt.Errorf("provider json.Unmarshal error: %s", err)
	// }

	// ArrOfStringIds := IntArrToStringArr(result)
	// return ArrOfStringIds, nil

	arrOfStringIds := IntArrToStringArr(mockAllItemIds)
	return arrOfStringIds, nil
}

func IntArrToStringArr(intArr []int) []string {
	var stringArr []string
	for _, i := range intArr {
		stringArr = append(stringArr, strconv.Itoa(i))
	}
	return stringArr
}

var mockAllItemIds = []int{
	24,
	33,
	46,
	56,
	57,
	58,
	59,
	60,
	61,
	62,
	63,
	64,
	65,
	68,
}
