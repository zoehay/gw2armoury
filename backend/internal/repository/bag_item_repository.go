package repository

import (
	"github.com/zoehay/gw2armoury/backend/internal/models"
	repositorymodels "github.com/zoehay/gw2armoury/backend/internal/repository/repository_models"
	"gorm.io/gorm"
)

type BagItemRepository interface {
	Create(BagItem *repositorymodels.GORMBagItem) (*repositorymodels.GORMBagItem, error)
	DeleteByCharacterName(characterName string) error
	GetByCharacterName(characterName string) ([]repositorymodels.GORMBagItem, error)
	GetIds() ([]int, error)
}

type GORMBagItemRepository struct {
	DB *gorm.DB
}

func NewGORMBagItemRepository(db *gorm.DB) GORMBagItemRepository {
	return GORMBagItemRepository{
		DB: db,
	}
}

func (repository *GORMBagItemRepository) Create(BagItem *repositorymodels.GORMBagItem) (*repositorymodels.GORMBagItem, error) {
	err := repository.DB.Create(&BagItem).Error
	if err != nil {
		return nil, err
	}

	return BagItem, nil
}

func (repository *GORMBagItemRepository) DeleteByCharacterName(characterName string) error {
	err := repository.DB.Where("character_name = ?", characterName).Delete(&repositorymodels.GORMBagItem{}).Error

	return err
}

func (repository *GORMBagItemRepository) GetByCharacterName(characterName string) ([]repositorymodels.GORMBagItem, error) {
	var bagItems []repositorymodels.GORMBagItem
	err := repository.DB.Where("character_name = ?", characterName).Find(&bagItems).Error
	if err != nil {
		return nil, err
	}

	return bagItems, nil
}

func (repository *GORMBagItemRepository) GetIds() ([]int, error) {
	var bagItemIds []int
	err := repository.DB.Model(&repositorymodels.GORMBagItem{}).Pluck("bag_item_id", &bagItemIds).Error
	if err != nil {
		return nil, err
	}

	return bagItemIds, nil

}

func (repository *GORMBagItemRepository) GetDetailsByCharacterName(characterName string) ([]models.BagItem, error) {
	var bagItemDetails []models.BagItem

	err := repository.DB.Table("gorm_bag_items").
		Select("gorm_bag_items.*, gorm_items.icon").
		Joins("left join gorm_items on gorm_bag_items.bag_item_id = gorm_items.id").
		Where("gorm_bag_items.character_name = ?", characterName).
		Scan(&bagItemDetails).Error

	if err != nil {
		return nil, err
	}

	return bagItemDetails, nil

}

func (repository *GORMBagItemRepository) GetDetailsByAccountID(accountID string) ([]models.BagItem, error) {
	var bagItemDetails []models.BagItem

	err := repository.DB.Table("gorm_bag_items").
		Select("gorm_bag_items.*, gorm_items.icon").
		Joins("left join gorm_items on gorm_bag_items.bag_item_id = gorm_items.id").
		Where("gorm_bag_items.account_id = ?", accountID).
		Scan(&bagItemDetails).Error

	if err != nil {
		return nil, err
	}

	return bagItemDetails, nil

}
