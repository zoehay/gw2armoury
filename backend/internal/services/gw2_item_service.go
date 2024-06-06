package services

import (
	"fmt"
	"strconv"

	"github.com/lib/pq"
	gw2api "github.com/zoehay/gw2armoury/backend/internal/gw2_api"
	apimodels "github.com/zoehay/gw2armoury/backend/internal/gw2_api/api_models"
	"github.com/zoehay/gw2armoury/backend/internal/repository"
)

type ItemServiceInterface interface {
	GetAndStoreItemsById(stringOfIds string) error
	GetAndStoreAllItems() error
}

type ItemService struct {
	gormItemRepository *repository.GORMItemRepository
}

func NewItemService(itemRepository *repository.GORMItemRepository) *ItemService {
	return &ItemService{
		gormItemRepository: itemRepository,
	}
}

func (service *ItemService) GetAndStoreItemsById(ids []int) error {
	apiItems, err := gw2api.GetItemsByIds(ids)
	if err != nil {
		return fmt.Errorf("service error using provider: %s", err)
	}

	for _, item := range apiItems {
		gormItem := apimodels.APIItemToGORMItem(item)
		_, err := service.gormItemRepository.Create(&gormItem)
		if err != nil {
			return fmt.Errorf("service error using gorm create: %s", err)
		}
	}
	return nil
}

func (service *ItemService) GetAndStoreAllItems() error {
	allItemIds, err := gw2api.GetAllItemIds()

	if err != nil {
		return fmt.Errorf("service error getting all itemIds: %s", err)
	}

	itemIdChunks := SplitArray(allItemIds, 3)
	fmt.Println(itemIdChunks)

	for _, idChunk := range itemIdChunks {
		err = service.GetAndStoreItemsById(idChunk)
		if err != nil {
			return fmt.Errorf("service error getting itemId chunk %d: %s", idChunk, err)
		}
	}

	return nil
}

func SplitArray(arr []int, chunkSize int) [][]int {
	var result [][]int

	for i := 0; i < len(arr); i += chunkSize {
		end := i + chunkSize
		if end > len(arr) {
			end = len(arr)
		}
		result = append(result, arr[i:end])
	}

	return result

}

func (service *ItemService) GetAndStoreEachByIds(itemIds []int) error {
	apiItems, err := gw2api.GetItemsByIds(itemIds)
	if err != nil {
		return fmt.Errorf("provider error requesting items: %s", err)
	}

	var duplicateKeyErrorIds []int
	for _, item := range apiItems {
		gormItem := apimodels.APIItemToGORMItem(item)
		_, err := service.gormItemRepository.Create(&gormItem)
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

func isDuplicateKeyError(err error) bool {
	if err, ok := err.(*pq.Error); ok {
		return err.Code == "23505"
	}
	return false
}

func IntArrToStringArr(intArr []int) []string {
	var stringArr []string
	for _, num := range intArr {
		stringArr = append(stringArr, strconv.Itoa(num))
	}
	return stringArr

}
