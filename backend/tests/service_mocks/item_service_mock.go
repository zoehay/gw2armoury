package servicemocks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/lib/pq"
	apimodels "github.com/zoehay/gw2armoury/backend/internal/gw2_api/api_models"
	"github.com/zoehay/gw2armoury/backend/internal/repository"
)

type ItemServiceMock struct {
	GormItemRepository *repository.GORMItemRepository
}

func NewItemServiceMock(itemRepository *repository.GORMItemRepository) *ItemServiceMock {
	return &ItemServiceMock{
		GormItemRepository: itemRepository,
	}
}

func (service *ItemServiceMock) GetAndStoreAllItems() error {
	apiItems, err := service.readItemFromFile("/Users/zoehay/Projects/gw2armoury/backend/test_data/item_test_data.txt")

	if err != nil {
		return fmt.Errorf("error reading from test data file: %s", err)
	}

	var duplicateKeyErrorIds []int
	for _, item := range apiItems {
		gormItem := apimodels.APIItemToGORMItem(item)
		_, err := service.GormItemRepository.Create(&gormItem)
		if err != nil {
			if isDuplicateKeyError(err) {
				duplicateKeyErrorIds = append(duplicateKeyErrorIds, int(item.ID))
			} else {
				return fmt.Errorf("gorm error adding item id %d: %s", item.ID, err)
			}
		}
	}

	if len(duplicateKeyErrorIds) != 0 {
		fmt.Printf("skipped adding duplicate values %#v\n", duplicateKeyErrorIds)
	}

	return nil

}

func (service *ItemServiceMock) readItemFromFile(filepath string) ([]apimodels.APIItem, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var items []apimodels.APIItem
	err = json.Unmarshal(content, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func isDuplicateKeyError(err error) bool {
	if err, ok := err.(*pq.Error); ok {
		return err.Code == "23505"
	}
	return false
}
