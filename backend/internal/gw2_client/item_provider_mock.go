package gw2api

import (
	"encoding/json"
	"fmt"
	"os"

	gw2models "github.com/zoehay/gw2armoury/backend/internal/gw2_client/gw2_models"
)

type ItemProviderMock struct{}

func (itemProvider *ItemProviderMock) GetItemsByIds(intArrIds []int) ([]gw2models.GW2Item, error) {
	apiItems, err := itemProvider.readItemFromFile("../test_data/item_test_data.txt")

	if err != nil {
		return nil, fmt.Errorf("error reading from test data file: %s", err)
	}

	return apiItems, nil
}

func (itemProvider *ItemProviderMock) GetAllItemIds() ([]int, error) {
	return mockAllItemIds, nil
}

func (service *ItemProviderMock) readItemFromFile(filepath string) ([]gw2models.GW2Item, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var items []gw2models.GW2Item
	err = json.Unmarshal(content, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

var mockAllItemIds = []int{
	24,
}
