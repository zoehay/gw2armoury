package repository

import (
	repositorymodels "github.com/zoehay/gw2armoury/backend/internal/repository/repository_models"
	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(Account *repositorymodels.Account) (*repositorymodels.Account, error)
	DeleteByCharacterName(characterName string) error
	GetByCharacterName(characterName string) ([]repositorymodels.Account, error)
	GetIds() ([]int, error)
}

type GORMAccountRepository struct {
	DB *gorm.DB
}

func NewGORMAccountRepository(db *gorm.DB) GORMAccountRepository {
	return GORMAccountRepository{
		DB: db,
	}
}

func (repository *GORMAccountRepository) GetBySession(sessionID string) (*repositorymodels.Account, error) {
	var account repositorymodels.Account

	err := repository.DB.Where("Session = ?", sessionID).First(&account).Error
	if err != nil {
		return nil, err
	}

	return &account, nil

}

func (repository *GORMAccountRepository) GetByName(accountName string) (*repositorymodels.Account, error) {
	var account repositorymodels.Account

	err := repository.DB.Where("AccountName = ?", accountName).First(&account).Error
	if err != nil {
		return nil, err
	}

	return &account, nil

}

func (repository *GORMAccountRepository) Create(account *repositorymodels.Account) (*repositorymodels.Account, error) {

	err := repository.DB.Create(&account).Error
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (repository *GORMAccountRepository) UpdateSession(accountID string, sessionID string) (*repositorymodels.Account, error) {
	var account repositorymodels.Account

	err := repository.DB.Model(&account).Where("AccountID = ?", accountID).Update("Session", sessionID).Error
	if err != nil {
		return nil, err
	}

	return &account, nil
}
