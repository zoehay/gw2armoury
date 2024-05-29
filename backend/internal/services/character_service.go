package services

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	gw2api "github.com/zoehay/gw2armoury/backend/internal/gw2_api"
	apimodels "github.com/zoehay/gw2armoury/backend/internal/gw2_api/api_models"
	"github.com/zoehay/gw2armoury/backend/internal/repository"
)

type CharacterServiceInterface interface {
	UpdateInventory() error
	GetAllCharacterNames() ([]string, error)
	UpdateCharacterInventory() error
}

type CharacterService struct {
	gormBagItemRepository *repository.GormBagItemRepository
}

func NewCharacterService(bagItemRepository *repository.GormBagItemRepository) *CharacterService {
	return &CharacterService{
		gormBagItemRepository: bagItemRepository,
	}
}

// func UpdateInventory(apiKey)
// 	getAllCharacterNames
// 	for character of characters
// 		UpdateCharacterInventory

// func GetAllCharacterNames(apiKey)

// func UpdateCharacterInventory(apiKey, characterName)
// 	newBagItems = provider.GetCharacterInventory(apiKey, characterName)
//
// 	BagItemRepository.DeleteBagItemsByCharacter(characterName)
// 	BagItemRespository.CreateCharacterBagItems()

func (service *CharacterService) GetAndStoreAllCharacters() error {
	////////////
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	apiKey := os.Getenv("TEST_API_KEY")
	////////////

	characters, err := gw2api.GetAllCharacters(apiKey)
	fmt.Println("got characters", characters)
	if err != nil {
		return fmt.Errorf("service error using provider: %s", err)
	}

	for _, character := range characters {
		fmt.Println(character.Name)
		err = service.ClearCharacterInventory(character)
		if err != nil {
			return err
		}
		err = service.StoreCharacterInventory(character)
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *CharacterService) StoreCharacterInventory(character apimodels.ApiCharacter) error {
	apiBags := character.Bags

	if apiBags != nil {
		for _, bag := range *apiBags {
			for _, bagItem := range bag.Inventory {
				if bagItem != nil {
					gormBagItem := apimodels.ApiBagToGormBagItem(character.Name, *bagItem)
					_, err := service.gormBagItemRepository.Create(&gormBagItem)
					if err != nil {
						return fmt.Errorf("service error using gorm create bagitem %d for character %s: %s", bagItem.Id, character.Name, err)
					}
				}
			}
		}
	}
	return nil
}

func (service *CharacterService) ClearCharacterInventory(character apimodels.ApiCharacter) error {
	err := service.gormBagItemRepository.DeleteByCharacterName(character.Name)
	if err != nil {
		return fmt.Errorf("service error using gorm delete bagitems for character %s: %s", character.Name, err)
	}

	return nil
}
