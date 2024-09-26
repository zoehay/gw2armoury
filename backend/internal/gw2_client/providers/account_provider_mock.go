package providers

import (
	"encoding/json"
	"fmt"
	"os"

	gw2models "github.com/zoehay/gw2armoury/backend/internal/gw2_client/models"
)

type AccountProviderMock struct{}

func (accountProvider *AccountProviderMock) GetAccount(apiKey string) (*gw2models.GW2Account, error) {
	apiAccount, err := accountProvider.ReadAccountFromFile("../../test_data/account_test_data.txt")

	if err != nil {
		return nil, fmt.Errorf("error reading from test data file: %s", err)
	}

	return apiAccount, nil
}

func (accountProvider *AccountProviderMock) ReadAccountFromFile(filepath string) (*gw2models.GW2Account, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var account gw2models.GW2Account
	err = json.Unmarshal(content, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
