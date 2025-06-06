package repositories

import (
	dbmodels "github.com/zoehay/gw2armoury/backend/internal/db/models"
	"gorm.io/gorm"
)

type BagItemRepositoryInterface interface {
	Create(BagItem *dbmodels.DBBagItem) (*dbmodels.DBBagItem, error)
	DeleteByAccountID(accountID string) error
	DeleteByCharacterName(characterName string) error
	GetByCharacterName(characterName string) ([]dbmodels.DBBagItem, error)
	DeleteSharedInventory(accountID string) error
	GetIds() ([]int, error)
	GetDetailBagItemByCharacterName(characterName string) ([]dbmodels.DBDetailBagItem, error)
	GetDetailBagItemByAccountID(accountID string) ([]dbmodels.DBDetailBagItem, error)
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

func (repository *BagItemRepository) DeleteByCharacterName(characterName string) error {
	err := repository.DB.Where("character_name = ?", characterName).Delete(&dbmodels.DBBagItem{}).Error

	return err
}

func (repository *BagItemRepository) GetByCharacterName(characterName string) ([]dbmodels.DBBagItem, error) {
	var bagItems []dbmodels.DBBagItem
	err := repository.DB.Where("character_name = ?", characterName).Find(&bagItems).Error
	if err != nil {
		return nil, err
	}

	return bagItems, nil
}

func (repository *BagItemRepository) DeleteSharedInventory(accountID string) error {
	err := repository.DB.Where("db_bag_items.account_id = ?", accountID).Where("character_name IS NULL OR character_name = ''").Delete(&dbmodels.DBBagItem{}).Error

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

func (repository *BagItemRepository) GetDetailBagItemByCharacterName(characterName string) ([]dbmodels.DBDetailBagItem, error) {
	var detailBagItems []dbmodels.DBDetailBagItem

	err := repository.DB.Table("db_bag_items").
		Select("db_bag_items.*, db_items.icon, db_items.name, db_items.description, db_items.rarity").
		Joins("left join db_items on db_bag_items.bag_item_id = db_items.id").
		Where("db_bag_items.character_name = ?", characterName).
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
