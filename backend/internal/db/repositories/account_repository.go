package repositories

import (
	dbmodels "github.com/zoehay/gw2armoury/backend/internal/db/models"
	"gorm.io/gorm"
)

type AccountRepositoryInterface interface {
	GetBySession(sessionID string) (*dbmodels.DBAccount, error)
	GetByID(id string) (*dbmodels.DBAccount, error)
	GetByName(name string) (*dbmodels.DBAccount, error)
	Create(account *dbmodels.DBAccount) (*dbmodels.DBAccount, error)
	UpdateSession(accountID string, sessionID string) (*dbmodels.DBAccount, error)
}

type AccountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return AccountRepository{
		DB: db,
	}
}

func (repository *AccountRepository) GetBySession(sessionID string) (*dbmodels.DBAccount, error) {
	var account dbmodels.DBAccount

	err := repository.DB.Where("Session = ?", sessionID).First(&account).Error
	if err != nil {
		return nil, err
	}

	return &account, nil

}

func (repository *AccountRepository) GetByID(id string) (*dbmodels.DBAccount, error) {
	var account dbmodels.DBAccount

	err := repository.DB.Where("account_id = ?", id).First(&account).Error
	if err != nil {
		return nil, err
	}

	return &account, nil

}

func (repository *AccountRepository) GetByName(name string) (*dbmodels.DBAccount, error) {
	var account dbmodels.DBAccount

	err := repository.DB.Where("AccountName = ?", name).First(&account).Error
	if err != nil {
		return nil, err
	}

	return &account, nil

}

func (repository *AccountRepository) Create(account *dbmodels.DBAccount) (*dbmodels.DBAccount, error) {

	err := repository.DB.Create(&account).Error
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (repository *AccountRepository) UpdateSession(accountID string, sessionID string) (*dbmodels.DBAccount, error) {
	var account dbmodels.DBAccount

	err := repository.DB.Model(&account).Where("AccountID = ?", accountID).Update("Session", sessionID).Error
	if err != nil {
		return nil, err
	}

	return &account, nil
}
