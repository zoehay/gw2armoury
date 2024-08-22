package repositories

import (
	dbmodels "github.com/zoehay/gw2armoury/backend/internal/db/models"
	"gorm.io/gorm"
)

type SessionRepositoryInterface interface {
	Create(session *dbmodels.DBSession) (*dbmodels.DBSession, error)
	Delete(sessionID string) error
	Get(sessionID string) (*dbmodels.DBSession, error)
	// Reset(session *repositorymodels.Session) (*repositorymodels.Session, error)
}

type SessionRepository struct {
	DB *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return SessionRepository{
		DB: db,
	}
}

func (repository *SessionRepository) Create(session *dbmodels.DBSession) (*dbmodels.DBSession, error) {
	err := repository.DB.Create(&session).Error
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (repository *SessionRepository) Delete(sessionID string) error {
	var session *dbmodels.DBSession
	err := repository.DB.Where("session_id = ?", sessionID).Delete(&session).Error

	return err
}

func (repository *SessionRepository) Get(sessionID string) (*dbmodels.DBSession, error) {
	var session *dbmodels.DBSession
	err := repository.DB.Where("session_id = ?", sessionID).Find(&session).Error
	if err != nil {
		return nil, err
	}

	return session, err
}
