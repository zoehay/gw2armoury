package repositories

import "gorm.io/gorm"

type Repository interface {
	Migrate() error
	GetAll() (interface{}, error)
	GetById(id int) (interface{}, error)
	Create(entity interface{}) (interface{}, error)
}

type GORMRepository struct {
	db *gorm.DB
}

func NewGORMRepository(db *gorm.DB) *GORMRepository {
	return &GORMRepository{
		db: db,
	}
}
