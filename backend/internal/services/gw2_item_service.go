package services

import (
	"fmt"
	"strings"

	gw2api "github.com/zoehay/gw2armoury/backend/internal/gw2_api"
	apimodels "github.com/zoehay/gw2armoury/backend/internal/gw2_api/api_models"
	"github.com/zoehay/gw2armoury/backend/internal/repository"
)

type ItemServiceInterface interface {
	GetAndStoreItemsById(stringOfIds string) error
	GetAndStoreAllItems() error
}

type ItemService struct {
	gormItemRepository *repository.GormItemRepository
}

func NewItemService(itemRepository *repository.GormItemRepository) *ItemService {
	return &ItemService{
		gormItemRepository: itemRepository,
	}
}

func (service *ItemService) GetAndStoreItemsById(idsString string) error {
	apiItems, err := gw2api.GetItemsById(idsString)
	if err != nil {
		return fmt.Errorf("service error using provider: %s", err)
	}

	for _, item := range apiItems {
		gormItem := apimodels.ApiItemToGormItem(item)
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
		idString := strings.Join(idChunk, ",")
		fmt.Println("get and store items", idString)
		err = service.GetAndStoreItemsById(idString)
		if err != nil {
			return fmt.Errorf("service error getting itemId chunk %s: %s", idString, err)
		}
	}

	return nil
}

func SplitArray(arr []string, chunkSize int) [][]string {
	var result [][]string

	for i := 0; i < len(arr); i += chunkSize {
		end := i + chunkSize
		if end > len(arr) {
			end = len(arr)
		}
		result = append(result, arr[i:end])
	}

	return result

}
