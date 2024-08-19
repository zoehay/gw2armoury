package services

import (
	"fmt"

	gw2client "github.com/zoehay/gw2armoury/backend/internal/gw2_client"
	"github.com/zoehay/gw2armoury/backend/internal/repository"
)

type AccountServiceInterface interface {
	GetAccountID(apiKey string) (*string, error)
}

type AccountService struct {
	GORMAccountRepository *repository.GORMAccountRepository
	AccountProvider       gw2client.AccountDataProvider
}

func NewAccountService(accountRepository *repository.GORMAccountRepository, accountProvider gw2client.AccountDataProvider) *AccountService {
	return &AccountService{
		GORMAccountRepository: accountRepository,
		AccountProvider:       accountProvider,
	}
}

func (service *AccountService) GetAccountID(apiKey string) (*string, error) {
	account, err := service.AccountProvider.GetAccount(apiKey)
	if err != nil {
		return nil, fmt.Errorf("service error using provider could not get account id: %s", err)
	}
	if account.ID == nil {
		return nil, fmt.Errorf("service error no account id: %s", err)
	}

	return account.ID, nil
}
