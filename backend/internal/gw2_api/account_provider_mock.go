package gw2api

import (
	"encoding/json"
	"fmt"
	"os"

	apimodels "github.com/zoehay/gw2armoury/backend/internal/gw2_api/api_models"
)

type AccountProviderMock struct{}

func (accountProvider *AccountProviderMock) GetAccount(apiKey string) (*apimodels.APIAccount, error) {
	apiAccount, err := accountProvider.ReadAccountFromFile("../test_data/account_test_data.txt")

	if err != nil {
		return nil, fmt.Errorf("error reading from test data file: %s", err)
	}

	return apiAccount, nil
}

func (accountProvider *AccountProviderMock) ReadAccountFromFile(filepath string) (*apimodels.APIAccount, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var account apimodels.APIAccount
	err = json.Unmarshal(content, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
