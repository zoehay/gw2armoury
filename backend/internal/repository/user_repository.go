package repository

import (
	repositorymodels "github.com/zoehay/gw2armoury/backend/internal/repository/repository_models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(User *repositorymodels.User) (*repositorymodels.User, error)
	DeleteByCharacterName(characterName string) error
	GetByCharacterName(characterName string) ([]repositorymodels.User, error)
	GetIds() ([]int, error)
}

type GORMUserRepository struct {
	DB *gorm.DB
}

func NewGORMUserRepository(db *gorm.DB) GORMUserRepository {
	return GORMUserRepository{
		DB: db,
	}
}

func (repository *GORMItemRepository) GetBySession(session gorm.Session) (*repositorymodels.User, error) {
	var user repositorymodels.User

	err := repository.DB.Where("Session = ?", session).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil

}
