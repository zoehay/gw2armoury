package services

import (
	"fmt"

	gw2api "github.com/zoehay/gw2armoury/backend/internal/gw2_api"
)

type AccountServiceInterface interface {
	GetAccountID(apiKey string) (*string, error)
}

type AccountService struct {
}

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (service *AccountService) GetAccountID(apiKey string) (*string, error) {
	account, err := gw2api.GetAccount(apiKey)
	if err != nil {
		return nil, fmt.Errorf("service error using provider could not get account id: %s", err)
	}
	if account.ID == nil {
		return nil, fmt.Errorf("service error no account id: %s", err)
	}

	return account.ID, nil
}
