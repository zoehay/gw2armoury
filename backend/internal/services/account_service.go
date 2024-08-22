package services

import (
	"fmt"

	"github.com/zoehay/gw2armoury/backend/internal/db/repository"
	"github.com/zoehay/gw2armoury/backend/internal/gw2_client/providers"
)

type AccountServiceInterface interface {
	GetAccountID(apiKey string) (*string, error)
}

type AccountService struct {
	AccountRepository *repository.AccountRepository
	AccountProvider   providers.AccountDataProvider
}

func NewAccountService(accountRepository *repository.AccountRepository, accountProvider providers.AccountDataProvider) *AccountService {
	return &AccountService{
		AccountRepository: accountRepository,
		AccountProvider:   accountProvider,
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
