package services

import (
	gw2api "github.com/zoehay/gw2armoury/backend/internal/gw2_api"
	apimodels "github.com/zoehay/gw2armoury/backend/internal/gw2_api/api_models"
	"github.com/zoehay/gw2armoury/backend/internal/repository"
	repositorymodels "github.com/zoehay/gw2armoury/backend/internal/repository/repository_models"
)

type ItemServiceInterface interface {
	GetAndStoreSomeDbItems() (repositorymodels.GormItem, error)
}

type ItemService struct {
	gormItemRepository *repository.GormItemRepository
}

func NewItemService() *ItemService {
	return &ItemService{}
}

func (service *ItemService) GetAndStoreSomeDbItems() error {
	apiItems, err := gw2api.GetSomeItems()
	if err != nil {
		return err
	}

	var gormItems []*repositorymodels.GormItem
	for _, item := range *apiItems {
		gormItem := apimodels.ApiItemToGormItem(item)
		gormItems = append(gormItems, &gormItem)
	}

	err = service.gormItemRepository.CreateMany(gormItems)

	if err != nil {
		return err
	}

	return nil

}
