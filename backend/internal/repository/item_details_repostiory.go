package repository

import (
	repositorymodels "github.com/zoehay/gw2armoury/backend/internal/repository/repository_models"
	"gorm.io/gorm"
)

type ItemDetailsRepository interface {
	GetBagItemDetailsByCharacterName(characterName string) ([]repositorymodels.GORMBagItem, error)
}

type GORMItemDetailsRepository struct {
	DB *gorm.DB
}

func NewGORMItemDetailsRepository(db *gorm.DB) GORMItemDetailsRepository {
	return GORMItemDetailsRepository{
		DB: db,
	}
}

func (repository *GORMItemDetailsRepository) GetBagItemDetailsByCharacterName(characterName string) ([]repositorymodels.GORMBagItem, error) {
	var bagItemDetails []repositorymodels.GORMBagItem

	err := repository.DB.Table("gorm_bag_item").
		Select("gorm_bag_item.*, gorm_item.*").
		Joins("left join gorm_bag_item on gorm_bag_item.id = item.id").
		Where("gorm_bag_item.name = ?", characterName).
		Scan(&bagItemDetails).Error

	if err != nil {
		return nil, err
	}

	return bagItemDetails, nil

}
