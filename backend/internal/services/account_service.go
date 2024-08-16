package services

import (
	"fmt"

	gw2api "github.com/zoehay/gw2armoury/backend/internal/gw2_api"
	"github.com/zoehay/gw2armoury/backend/internal/repository"
)

type AccountServiceInterface interface {
	GetAccountID(apiKey string) (*string, error)
}

type AccountService struct {
	GORMAccountRepository *repository.GORMAccountRepository
	AccountProvider       gw2api.AccountDataProvider
}

func NewAccountService(accountRepository *repository.GORMAccountRepository, accountProvider gw2api.AccountDataProvider) *AccountService {
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
