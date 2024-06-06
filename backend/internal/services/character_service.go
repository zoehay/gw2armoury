package services

import (
	"fmt"

	gw2api "github.com/zoehay/gw2armoury/backend/internal/gw2_api"
	apimodels "github.com/zoehay/gw2armoury/backend/internal/gw2_api/api_models"
	"github.com/zoehay/gw2armoury/backend/internal/repository"
)

type CharacterServiceInterface interface {
	GetAndStoreAllCharacters(apiKey string) error
	storeCharacterInventory(character apimodels.APICharacter) error
	clearCharacterInventory(character apimodels.APICharacter) error
}

type CharacterService struct {
	gormBagItemRepository *repository.GORMBagItemRepository
}

func NewCharacterService(bagItemRepository *repository.GORMBagItemRepository) *CharacterService {
	return &CharacterService{
		gormBagItemRepository: bagItemRepository,
	}
}

func (service *CharacterService) GetAndStoreAllCharacters(apiKey string) error {
	characters, err := gw2api.GetAllCharacters(apiKey)
	if err != nil {
		return fmt.Errorf("service error using provider: %s", err)
	}

	tx := service.gormBagItemRepository.DB.Begin()

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, character := range characters {
		fmt.Println(character.Name)

		err = service.clearCharacterInventory(character)
		if err != nil {
			return err
		}
		err = service.storeCharacterInventory(character)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error

}

func (service *CharacterService) storeCharacterInventory(character apimodels.APICharacter) error {
	apiBags := character.Bags

	if apiBags != nil {
		for _, bag := range *apiBags {
			for _, bagItem := range bag.Inventory {
				if bagItem != nil {
					gormBagItem := apimodels.APIBagToGORMBagItem(character.Name, *bagItem)
					_, err := service.gormBagItemRepository.Create(&gormBagItem)
					if err != nil {
						return fmt.Errorf("service error using gorm create bagitem %d for character %s: %s", bagItem.ID, character.Name, err)
					}
				}
			}
		}
	}
	return nil
}

func (service *CharacterService) clearCharacterInventory(character apimodels.APICharacter) error {
	err := service.gormBagItemRepository.DeleteByCharacterName(character.Name)
	if err != nil {
		return fmt.Errorf("service error using gorm delete bagitems for character %s: %s", character.Name, err)
	}

	return nil
}
