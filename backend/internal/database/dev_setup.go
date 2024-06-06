package database

import (
	"fmt"

	"github.com/zoehay/gw2armoury/backend/internal/repository"
	"github.com/zoehay/gw2armoury/backend/internal/services"
)

func GetItemsInCharacterBags(itemService *services.ItemService, bagItemRepository repository.GORMBagItemRepository) error {
	itemIds, err := bagItemRepository.GetIds()
	fmt.Println("number of items", len(itemIds))
	if err != nil {
		return err
	}

	if len(itemIds) > 500 {
		return fmt.Errorf("more than 500 items, maybe wait")
	}

	errors := itemService.GetAndStoreByIds(itemIds)
	if len(errors) != 0 {
		fmt.Println(errors)
		return fmt.Errorf("encountered errors getting and storing items")
	}

	return nil
}
