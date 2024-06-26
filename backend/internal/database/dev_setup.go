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
		return fmt.Errorf("dev setup more than 500 items, maybe wait")
	}

	err = itemService.GetAndStoreEachByIds(itemIds)
	if err != nil {
		return fmt.Errorf("dev setup error getting and storing items : %s", err)
	}

	return nil
}
