package services

import (
	"fmt"

	"github.com/zoehay/gw2armoury/backend/internal/db/repositories"
	gw2models "github.com/zoehay/gw2armoury/backend/internal/gw2_client/models"
	"github.com/zoehay/gw2armoury/backend/internal/gw2_client/providers"
)

type AccountServiceInterface interface {
	GetAccount(apiKey string) (*gw2models.GW2Account, error)
}

type AccountService struct {
	AccountRepository *repositories.AccountRepository
	AccountProvider   providers.AccountDataProvider
}

func NewAccountService(accountRepository *repositories.AccountRepository, accountProvider providers.AccountDataProvider) *AccountService {
	return &AccountService{
		AccountRepository: accountRepository,
		AccountProvider:   accountProvider,
	}
}

func (service *AccountService) GetAccount(apiKey string) (*gw2models.GW2Account, error) {
	account, err := service.AccountProvider.GetAccount(apiKey)
	if err != nil {
		return nil, fmt.Errorf("service error using provider could not get account id: %s", err)
	}
	if account.ID == nil {
		return nil, fmt.Errorf("service error no account id: %s", err)
	}

	return account, nil
}
