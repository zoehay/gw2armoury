package repository

import (
	repositorymodels "github.com/zoehay/gw2armoury/backend/internal/repository/repository_models"
	"gorm.io/gorm"
)

type ItemRepository interface {
	Create(item *repositorymodels.GORMItem) error
	CreateMany(items []*repositorymodels.GORMItem) error
	GetAll() ([]repositorymodels.GORMItem, error)
	GetFirst() (repositorymodels.GORMItem, error)
	GetById(id int) (repositorymodels.GORMItem, error)
}

type GORMItemRepository struct {
	DB *gorm.DB
}

func NewGORMItemRepository(db *gorm.DB) GORMItemRepository {
	return GORMItemRepository{
		DB: db,
	}
}

func (repository *GORMItemRepository) Create(item *repositorymodels.GORMItem) (*repositorymodels.GORMItem, error) {

	err := repository.DB.Create(&item).Error
	if err != nil {
		return nil, err
	}

	return item, nil

}

func (repository *GORMItemRepository) CreateMany(items []*repositorymodels.GORMItem) error {

	err := repository.DB.Create(&items).Error
	if err != nil {
		return err
	}

	return nil

}

func (repository *GORMItemRepository) GetAll() ([]repositorymodels.GORMItem, error) {
	var items []repositorymodels.GORMItem

	err := repository.DB.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil

}

func (repository *GORMItemRepository) GetFirst() (*repositorymodels.GORMItem, error) {
	var item repositorymodels.GORMItem

	err := repository.DB.First(&repositorymodels.GORMItem{}).Error
	if err != nil {
		return nil, err
	}

	return &item, nil

}

func (repository *GORMItemRepository) GetById(id int) (*repositorymodels.GORMItem, error) {
	var item repositorymodels.GORMItem

	err := repository.DB.First(&item, id).Error
	if err != nil {
		return nil, err
	}

	return &item, nil

}
