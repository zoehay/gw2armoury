package repositories

import (
	dbmodels "github.com/zoehay/gw2armoury/backend/internal/db/models"
	"gorm.io/gorm"
)

type BagItemRepositoryInterface interface {
	Create(BagItem *dbmodels.DBBagItem) (*dbmodels.DBBagItem, error)
	DeleteByCharacterName(characterName string) error
	GetByCharacterName(characterName string) ([]dbmodels.DBBagItem, error)
	GetIds() ([]int, error)
	GetIconBagItemByCharacterName(characterName string) ([]dbmodels.DBIconBagItem, error)
	GetIconBagItemByAccountID(accountID string) ([]dbmodels.DBIconBagItem, error)
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

func (repository *BagItemRepository) GetIds() ([]int, error) {
	var bagItemIds []int
	err := repository.DB.Model(&dbmodels.DBBagItem{}).Pluck("bag_item_id", &bagItemIds).Error
	if err != nil {
		return nil, err
	}

	return bagItemIds, nil

}

func (repository *BagItemRepository) GetIconBagItemByCharacterName(characterName string) ([]dbmodels.DBIconBagItem, error) {
	var bagItemWithIcon []dbmodels.DBIconBagItem

	err := repository.DB.Table("db_bag_items").
		Select("db_bag_items.*, db_items.icon, db_items.name").
		Joins("left join db_items on db_bag_items.bag_item_id = db_items.id").
		Where("db_bag_items.character_name = ?", characterName).
		Scan(&bagItemWithIcon).Error

	if err != nil {
		return nil, err
	}

	for i := range bagItemWithIcon {
		bagItemWithIcon[i].ToBagItem()
	}

	return bagItemWithIcon, nil

}

func (repository *BagItemRepository) GetIconBagItemByAccountID(accountID string) ([]dbmodels.DBIconBagItem, error) {
	var bagItemWithIcon []dbmodels.DBIconBagItem

	err := repository.DB.Table("db_bag_items").
		Select("db_bag_items.*, db_items.icon, db_items.name").
		Joins("left join db_items on db_bag_items.bag_item_id = db_items.id").
		Where("db_bag_items.account_id = ?", accountID).
		Scan(&bagItemWithIcon).Error

	if err != nil {
		return nil, err
	}

	for i := range bagItemWithIcon {
		bagItemWithIcon[i].ToBagItem()
	}

	return bagItemWithIcon, nil

}
