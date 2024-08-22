package services

import (
	"fmt"

	"github.com/zoehay/gw2armoury/backend/internal/db/repository"
	gw2models "github.com/zoehay/gw2armoury/backend/internal/gw2_client/models"
	"github.com/zoehay/gw2armoury/backend/internal/gw2_client/providers"
)

type CharacterServiceInterface interface {
	GetAndStoreAllCharacters(apiKey string) error
	storeCharacterInventory(character gw2models.GW2Character) error
	clearCharacterInventory(character gw2models.GW2Character) error
}

type CharacterService struct {
	BagItemRepository *repository.BagItemRepository
	CharacterProvider providers.CharacterDataProvider
}

func NewCharacterService(bagItemRepository *repository.BagItemRepository, characterProvider providers.CharacterDataProvider) *CharacterService {
	return &CharacterService{
		BagItemRepository: bagItemRepository,
		CharacterProvider: characterProvider,
	}
}

func (service *CharacterService) GetAndStoreAllCharacters(accountID string, apiKey string) error {
	characters, err := service.CharacterProvider.GetAllCharacters(apiKey)
	if err != nil {
		return fmt.Errorf("service error using provider could not get characters: %s", err)
	}

	tx := service.BagItemRepository.DB.Begin()

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
			tx.Rollback()
			return err
		}
		err = service.storeCharacterInventory(accountID, character)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error

}

func (service *CharacterService) storeCharacterInventory(accountID string, character gw2models.GW2Character) error {
	apiBags := character.Bags

	if apiBags != nil {
		for _, bag := range *apiBags {
			for _, bagItem := range bag.Inventory {
				if bagItem != nil {
					dbBagItem := bagItem.ToDBBagItem(accountID, character.Name)
					_, err := service.BagItemRepository.Create(&dbBagItem)
					if err != nil {
						return fmt.Errorf("service error using gorm create bagitem %d for character %s: %s", bagItem.ID, character.Name, err)
					}
				}
			}
		}
	}
	return nil
}

func (service *CharacterService) clearCharacterInventory(character gw2models.GW2Character) error {
	err := service.BagItemRepository.DeleteByCharacterName(character.Name)
	if err != nil {
		return fmt.Errorf("service error using gorm delete bagitems for character %s: %s", character.Name, err)
	}

	return nil
}
