package repository

import (
	repositorymodels "github.com/zoehay/gw2armoury/backend/internal/database/repository_models"
	"gorm.io/gorm"
)

type ItemRepositoryInterface interface {
	Create(item *repositorymodels.DBItem) (*repositorymodels.DBItem, error)
	CreateMany(items []*repositorymodels.DBItem) error
	GetAll() ([]repositorymodels.DBItem, error)
	GetFirst() (*repositorymodels.DBItem, error)
	GetById(id int) (*repositorymodels.DBItem, error)
}

type ItemRepository struct {
	DB *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return ItemRepository{
		DB: db,
	}
}

func (repository *ItemRepository) Create(item *repositorymodels.DBItem) (*repositorymodels.DBItem, error) {

	err := repository.DB.Create(&item).Error
	if err != nil {
		return nil, err
	}

	return item, nil

}

func (repository *ItemRepository) CreateMany(items []*repositorymodels.DBItem) error {

	err := repository.DB.Create(&items).Error
	if err != nil {
		return err
	}

	return nil

}

func (repository *ItemRepository) GetAll() ([]repositorymodels.DBItem, error) {
	var items []repositorymodels.DBItem

	err := repository.DB.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil

}

func (repository *ItemRepository) GetFirst() (*repositorymodels.DBItem, error) {
	var item repositorymodels.DBItem

	err := repository.DB.First(&repositorymodels.DBItem{}).Error
	if err != nil {
		return nil, err
	}

	return &item, nil

}

func (repository *ItemRepository) GetById(id int) (*repositorymodels.DBItem, error) {
	var item repositorymodels.DBItem

	err := repository.DB.First(&item, id).Error
	if err != nil {
		return nil, err
	}

	return &item, nil

}

func (repository *ItemRepository) GetByIds(ids []int) (*[]repositorymodels.DBItem, error) {
	var items []repositorymodels.DBItem

	err := repository.DB.Where(ids).Find(&items).Error
	if err != nil {
		return nil, err
	}

	return &items, nil

}
