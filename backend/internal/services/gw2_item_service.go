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
	apiItems, err := gw2api.GetSomeItems()
	if err != nil {
		return err
	}

	// var gormItems []*repositorymodels.GormItem
	for _, item := range apiItems {
		fmt.Println("service", item.Name, item.ChatLink, item.Description)
		gormItem := apimodels.ApiItemToGormItem(item)
		fmt.Println("service convert", item.Name, item.ChatLink, item.Description)
		_, err := service.gormItemRepository.Create(&gormItem)
		if err != nil {
			return err
		}
		// gormItems = append(gormItems, &gormItem)
	}
	// err = service.gormItemRepository.CreateMany(gormItems)

	// if err != nil {
	// 	return err
	// }

	return nil

}
