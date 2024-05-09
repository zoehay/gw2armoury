package repository

import (
	repositorymodels "github.com/zoehay/gw2armoury/backend/internal/repository/repository_models"
	"gorm.io/gorm"
)

type ItemRepository interface {
	Create(item *repositorymodels.GormItem) error
	GetAll() ([]repositorymodels.GormItem, error)
	GetFirst() (repositorymodels.GormItem, error)
}

type GormItemRepository struct {
	DB *gorm.DB
}

func NewGormItemRepository(db *gorm.DB) GormItemRepository {
	return GormItemRepository{
		DB: db,
	}
}

func (repository *GormItemRepository) Create(item *repositorymodels.GormItem) (*repositorymodels.GormItem, error) {

	err := repository.DB.Create(&item).Error
	if err != nil {
		return nil, err
	}

	return item, nil

}

func (repository *GormItemRepository) CreateMany(items []*repositorymodels.GormItem) error {

	err := repository.DB.Create(items).Error
	if err != nil {
		return err
	}

	return nil

}

func (repository *GormItemRepository) GetAll() ([]repositorymodels.GormItem, error) {
	var items []repositorymodels.GormItem

	err := repository.DB.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil

}

func (repository *GormItemRepository) GetFirst() (*repositorymodels.GormItem, error) {
	var item repositorymodels.GormItem

	err := repository.DB.First(&repositorymodels.GormItem{}).Error
	if err != nil {
		return nil, err
	}

	return &item, nil

}

func (repository *GormItemRepository) GetById(id string) (*repositorymodels.GormItem, error) {
	var item repositorymodels.GormItem

	err := repository.DB.First(&repositorymodels.GormItem{}, id).Error
	if err != nil {
		return nil, err
	}

	return &item, nil

}
