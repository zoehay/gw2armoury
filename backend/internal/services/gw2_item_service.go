package services

import (
	"fmt"

	gw2api "github.com/zoehay/gw2armoury/backend/internal/gw2_api"
	apimodels "github.com/zoehay/gw2armoury/backend/internal/gw2_api/api_models"
	"github.com/zoehay/gw2armoury/backend/internal/repository"
)

type ItemServiceInterface interface {
	GetAndStoreSomeDbItems() error
}

type ItemService struct {
	gormItemRepository *repository.GormItemRepository
}

func NewItemService(itemRepository *repository.GormItemRepository) *ItemService {
	return &ItemService{
		gormItemRepository: itemRepository,
	}
}

func (service *ItemService) GetAndStoreSomeDbItems() error {
	apiItems, err := gw2api.GetSomeItems("24,68")
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

var allItemIds = []int{
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
	69,
	70,
	71,
	72,
	73,
	74,
	75,
	76,
	77,
	78,
	79,
	80,
}
