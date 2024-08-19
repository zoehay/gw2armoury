package repository

import (
	repositorymodels "github.com/zoehay/gw2armoury/backend/internal/database/repository_models"
	"gorm.io/gorm"
)

type AccountRepositoryInterface interface {
	GetBySession(sessionID string) (*repositorymodels.DBAccount, error)
	GetByName(accountName string) (*repositorymodels.DBAccount, error)
	Create(account *repositorymodels.DBAccount) (*repositorymodels.DBAccount, error)
	UpdateSession(accountID string, sessionID string) (*repositorymodels.DBAccount, error)
}

type AccountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return AccountRepository{
		DB: db,
	}
}

func (repository *AccountRepository) GetBySession(sessionID string) (*repositorymodels.DBAccount, error) {
	var account repositorymodels.DBAccount

	err := repository.DB.Where("Session = ?", sessionID).First(&account).Error
	if err != nil {
		return nil, err
	}

	return &account, nil

}

func (repository *AccountRepository) GetByName(accountName string) (*repositorymodels.DBAccount, error) {
	var account repositorymodels.DBAccount

	err := repository.DB.Where("AccountName = ?", accountName).First(&account).Error
	if err != nil {
		return nil, err
	}

	return &account, nil

}

func (repository *AccountRepository) Create(account *repositorymodels.DBAccount) (*repositorymodels.DBAccount, error) {

	err := repository.DB.Create(&account).Error
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (repository *AccountRepository) UpdateSession(accountID string, sessionID string) (*repositorymodels.DBAccount, error) {
	var account repositorymodels.DBAccount

	err := repository.DB.Model(&account).Where("AccountID = ?", accountID).Update("Session", sessionID).Error
	if err != nil {
		return nil, err
	}

	return &account, nil
}
