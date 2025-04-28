package services

import (
	"fmt"

	"github.com/zoehay/gw2armoury/backend/internal/db/repositories"
	gw2models "github.com/zoehay/gw2armoury/backend/internal/gw2_client/models"
	"github.com/zoehay/gw2armoury/backend/internal/gw2_client/providers"
)

type BagItemServiceInterface interface {
	GetAndStoreAllCharacters(accountID string, apiKey string) error
	storeCharacterInventory(accountID string, character gw2models.GW2Character) error
	clearCharacterInventory(character gw2models.GW2Character) error
	GetAndStoreAccountInventory(accountID string, apiKey string) error
}

type BagItemService struct {
	BagItemRepository *repositories.BagItemRepository
	CharacterProvider providers.CharacterDataProvider
	AccountProvider   providers.AccountDataProvider
}

func NewBagItemService(bagItemRepository *repositories.BagItemRepository, characterProvider providers.CharacterDataProvider, accountProvider providers.AccountDataProvider) *BagItemService {
	return &BagItemService{
		BagItemRepository: bagItemRepository,
		CharacterProvider: characterProvider,
		AccountProvider:   accountProvider,
	}
}

func (service *BagItemService) GetAndStoreAccountInventory(accountID string, apiKey string) error {

	accountInventory, err := service.AccountProvider.GetAccountInventory(apiKey)
	if err != nil {
		return fmt.Errorf("service error using provider could not get account inventory: %s", err)
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

	err = service.clearAccountInventory(accountID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = service.storeAccountInventory(accountID, accountInventory)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error

}

func (service *BagItemService) GetAndStoreAllCharacters(accountID string, apiKey string) error {
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

func (service *BagItemService) storeCharacterInventory(accountID string, character gw2models.GW2Character) error {
	apiBags := character.Bags

	if apiBags != nil {
		for _, bag := range *apiBags {
			for _, bagItem := range bag.Inventory {
				if bagItem != nil {
					dbBagItem := bagItem.ToDBBagItem(accountID, &character.Name)
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

func (service *BagItemService) storeAccountInventory(accountID string, accountInventory *[]gw2models.GW2BagItem) error {
	for _, bagItem := range *accountInventory {
		dbBagItem := bagItem.ToDBBagItem(accountID, nil)
		_, err := service.BagItemRepository.Create(&dbBagItem)
		if err != nil {
			return fmt.Errorf("service error using gorm create bagitem %d for account %s: %s", bagItem.ID, accountID, err)
		}
	}
	return nil
}

func (service *BagItemService) clearCharacterInventory(character gw2models.GW2Character) error {
	err := service.BagItemRepository.DeleteByCharacterName(character.Name)
	if err != nil {
		return fmt.Errorf("service error using gorm delete bagitems for character %s: %s", character.Name, err)
	}

	return nil
}

func (service *BagItemService) clearAccountInventory(accountID string) error {
	err := service.BagItemRepository.DeleteAccountInventory(accountID)
	if err != nil {
		return fmt.Errorf("service error using gorm delete bagitems for account %s: %s", accountID, err)
	}

	return nil
}
