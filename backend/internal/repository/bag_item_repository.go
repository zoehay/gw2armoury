package repository

import (
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

func (repostiory *GORMBagItemRepository) DeleteByCharacterName(characterName string) error {
	err := repostiory.DB.Where("character_name = ?", characterName).Delete(&repositorymodels.GORMBagItem{}).Error

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
