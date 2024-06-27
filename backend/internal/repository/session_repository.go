package repository

import (
	repositorymodels "github.com/zoehay/gw2armoury/backend/internal/repository/repository_models"
	"gorm.io/gorm"
)

type SessionRepository interface {
	Create(Session *repositorymodels.Session) (*repositorymodels.Session, error)
}

type GORMSessionRepository struct {
	DB *gorm.DB
}

func NewGORMSessionRepository(db *gorm.DB) GORMSessionRepository {
	return GORMSessionRepository{
		DB: db,
	}
}

func (repository *GORMSessionRepository) Create(session *repositorymodels.Session) (*repositorymodels.Session, error) {

	err := repository.DB.Create(&session).Error
	if err != nil {
		return nil, err
	}

	return session, nil
}
