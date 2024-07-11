package repository

import (
	repositorymodels "github.com/zoehay/gw2armoury/backend/internal/repository/repository_models"
	"gorm.io/gorm"
)

type SessionRepository interface {
	Create(session *repositorymodels.Session) (*repositorymodels.Session, error)
	Delete(session *repositorymodels.Session) error
	Get(sessionID string) (*repositorymodels.Session, error)
	Reset(session *repositorymodels.Session) (*repositorymodels.Session, error)
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

func (repository *GORMSessionRepository) Delete(sessionID string) error {
	var session *repositorymodels.Session
	err := repository.DB.Where("session_id = ?", sessionID).Delete(&session).Error

	return err
}

func (repository *GORMSessionRepository) Get(sessionID string) error {
	var session *repositorymodels.Session
	err := repository.DB.Where("session_id = ?", sessionID).Find(&session).Error

	return err
}
