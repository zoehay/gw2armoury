package repository

import "gorm.io/gorm"


type Repository interface {
	Migrate() error
	GetAll() (interface{}, error)
	GetById(id int) (interface{}, error)
	Create(entity interface{}) (interface{}, error)
}

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db:db,
	}
}

