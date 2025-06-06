package repositories

import (
	"github.com/zoehay/gw2armoury/backend/internal/api/models"
	dbmodels "github.com/zoehay/gw2armoury/backend/internal/db/models"
	"gorm.io/gorm"
)

type BagItemRepositoryInterface interface {
	Create(BagItem *dbmodels.DBBagItem) (*dbmodels.DBBagItem, error)
	DeleteByAccountID(accountID string) error
	DeleteByCharacterName(accountID string, characterName string) error
	GetByCharacterName(accountID string, characterName string) ([]dbmodels.DBBagItem, error)
	DeleteSharedInventory(accountID string) error
	GetIds() ([]int, error)
	GetDetailBagItemByCharacterName(accountID string, characterName string) ([]dbmodels.DBDetailBagItem, error)
	GetDetailBagItemByAccountID(accountID string) ([]dbmodels.DBDetailBagItem, error)
	GetAccountInventory(accountID string) (*models.AccountInventory, error)
}

type BagItemRepository struct {
	DB *gorm.DB
}

func NewBagItemRepository(db *gorm.DB) BagItemRepository {
	return BagItemRepository{
		DB: db,
	}
}

func (repository *BagItemRepository) Create(BagItem *dbmodels.DBBagItem) (*dbmodels.DBBagItem, error) {
	err := repository.DB.Create(&BagItem).Error
	if err != nil {
		return nil, err
	}

	return BagItem, nil
}

func (repository *BagItemRepository) DeleteByAccountID(accountID string) error {
	err := repository.DB.Where("db_bag_items.account_id = ?", accountID).Delete(&dbmodels.DBBagItem{}).Error

	return err
}

func (repository *BagItemRepository) DeleteByCharacterName(accountID string, characterName string) error {
	err := repository.DB.Where("account_id = ? AND character_name = ?", accountID, characterName).Delete(&dbmodels.DBBagItem{}).Error

	return err
}

func (repository *BagItemRepository) GetByCharacterName(accountID string, characterName string) ([]dbmodels.DBBagItem, error) {
	var bagItems []dbmodels.DBBagItem
	err := repository.DB.Where("account_id = ? AND character_name = ?", accountID, characterName).Find(&bagItems).Error
	if err != nil {
		return nil, err
	}

	return bagItems, nil
}

func (repository *BagItemRepository) DeleteSharedInventory(accountID string) error {
	err := repository.DB.Where("db_bag_items.account_id = ? AND character_name = ?", accountID, "Shared Inventory").Delete(&dbmodels.DBBagItem{}).Error

	return err
}

func (repository *BagItemRepository) GetIds() ([]int, error) {
	var bagItemIds []int
	err := repository.DB.Model(&dbmodels.DBBagItem{}).Pluck("bag_item_id", &bagItemIds).Error
	if err != nil {
		return nil, err
	}

	return bagItemIds, nil

}

func (repository *BagItemRepository) GetDetailBagItemByCharacterName(accountID string, characterName string) ([]dbmodels.DBDetailBagItem, error) {
	var detailBagItems []dbmodels.DBDetailBagItem

	err := repository.DB.Table("db_bag_items").
		Select("db_bag_items.*, db_items.icon, db_items.name, db_items.description, db_items.rarity").
		Joins("left join db_items on db_bag_items.bag_item_id = db_items.id").
		Where("db_bag_items.account_id = ? AND db_bag_items.character_name = ?", accountID, characterName).
		Scan(&detailBagItems).Error

	if err != nil {
		return nil, err
	}

	for i := range detailBagItems {
		detailBagItems[i].ToBagItem()
	}

	return detailBagItems, nil

}

func (repository *BagItemRepository) GetDetailBagItemByAccountID(accountID string) ([]dbmodels.DBDetailBagItem, error) {
	var detailBagItems []dbmodels.DBDetailBagItem

	err := repository.DB.Table("db_bag_items").
		Select("db_bag_items.*, db_items.icon, db_items.name, db_items.description, db_items.rarity").
		Joins("left join db_items on db_bag_items.bag_item_id = db_items.id").
		Where("db_bag_items.account_id = ?", accountID).
		Scan(&detailBagItems).Error

	if err != nil {
		return nil, err
	}

	for i := range detailBagItems {
		detailBagItems[i].ToBagItem()
	}

	return detailBagItems, nil

}

func (repository *BagItemRepository) GetAccountInventory(accountID string) (*models.AccountInventory, error) {
	var detailBagItems []dbmodels.DBDetailBagItem

	err := repository.DB.Table("db_bag_items").
		Select("db_bag_items.*, db_items.icon, db_items.name, db_items.description, db_items.rarity").
		Joins("left join db_items on db_bag_items.bag_item_id = db_items.id").
		Where("db_bag_items.account_id = ?", accountID).
		Scan(&detailBagItems).Error

	if err != nil {
		return nil, err
	}

	characterNameMap := map[string]models.Character{}
	var sharedInventory []models.BagItem
	var characters []models.Character

	for i := range detailBagItems {
		item := detailBagItems[i].ToBagItem()
		name := item.CharacterName

		if name == "Shared Inventory" {
			sharedInventory = append(sharedInventory, item)
		} else {
			entry, ok := characterNameMap[name]
			isEquipment := repository.isEquipment(item)

			if ok {
				if isEquipment {
					entry.Equipment = append(entry.Equipment, item)
					characterNameMap[name] = entry
				} else {
					entry.Inventory = append(entry.Inventory, item)
					characterNameMap[name] = entry
				}
			} else {
				newCharacter := &models.Character{
					Name:      name,
					Equipment: []models.BagItem{},
					Inventory: []models.BagItem{},
				}
				if isEquipment {
					newCharacter.Equipment = append(newCharacter.Equipment, item)
				} else {
					newCharacter.Inventory = append(newCharacter.Inventory, item)
				}
				characterNameMap[name] = *newCharacter
			}
		}

	}

	for character := range characterNameMap {
		characters = append(characters, characterNameMap[character])
	}

	var accountInventory models.AccountInventory
	accountInventory.AccountID = accountID
	accountInventory.SharedInventory = &sharedInventory
	accountInventory.Characters = &characters

	return &accountInventory, nil

}

func (repository *BagItemRepository) isEquipment(item models.BagItem) bool {
	if item.Slot != nil && *item.Slot != "" {
		return true
	} else {
		return false
	}
}
