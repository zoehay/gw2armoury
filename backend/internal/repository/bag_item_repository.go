package repository

import (
	repositorymodels "github.com/zoehay/gw2armoury/backend/internal/repository/repository_models"
	"gorm.io/gorm"
)

type BagItemRepository interface {
	DeleteBagItemsByCharacter()
	CreateCharacterBagItems()
}

type GormBagItemRepository struct {
	DB *gorm.DB
}

func NewGormBagItemRepository(db *gorm.DB) GormBagItemRepository {
	return GormBagItemRepository{
		DB: db,
	}
}

func (repository *GormBagItemRepository) Create(BagItem *repositorymodels.GormBagItem) (*repositorymodels.GormBagItem, error) {

	err := repository.DB.Create(&BagItem).Error
	if err != nil {
		return nil, err
	}

	return BagItem, nil

}

func (repostiory *GormBagItemRepository) DeleteBagItemsByCharacterName(characterName string) error {
	err := repostiory.DB.Where("character_name = ?", characterName).Delete(&repositorymodels.GormBagItem{}).Error
	return err
}

// func (repository *GormBagItemRepository) GetBagItemsByCharacterName(characterName string) (*repositorymodels.GormBagItem, error) {
// 	var BagItem repositorymodels.GormBagItem

// 	err := repository.DB.First(&BagItem, id).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &BagItem, nil

// }
