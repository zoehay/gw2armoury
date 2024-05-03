package repository

import (
	"github.com/zoehay/gw2armoury/backend/internal/models"
	"gorm.io/gorm"
)

type ItemRepository interface {
	Create(item *models.Item) error
	GetAll() ([]models.Item, error) 
	GetFirst() (models.Item, error)
}

type GormItemRepository struct {
	DB *gorm.DB
}

func NewGormItemRepository(db *gorm.DB) GormItemRepository{
	return GormItemRepository{
		DB: db,
	}
}

func (repository *GormItemRepository) Create(item *models.Item) (*models.Item, error) {

	err := repository.DB.Create(item).Error
	if err != nil {
		return nil, err
	}

	return item, nil
	
}

func (repository *GormItemRepository) GetAll() ([]models.Item, error) {
	var items []models.Item

	err := repository.DB.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil

}

func (repository *GormItemRepository) GetFirst() (*models.Item, error) {
	var item models.Item

	err := repository.DB.First(&models.Item{}).Error
	if err != nil {
		return nil, err
	}

	return &item, nil

}

func (repository *GormItemRepository) GetById(id string) (*models.Item, error) {
	var item models.Item

	err := repository.DB.First(&models.Item{}, id).Error
	if err != nil {
		return nil, err
	}

	return &item, nil

}