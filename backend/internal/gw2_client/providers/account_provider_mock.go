package providers

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	gw2models "github.com/zoehay/gw2armoury/backend/internal/gw2_client/models"
)

type AccountProviderMock struct{}

func (accountProvider *AccountProviderMock) GetAccount(apiKey string) (*gw2models.GW2Account, error) {
	wd, _ := os.Getwd()
	isTesting := strings.Contains(wd, "test")
	leadingFilepath := ""

	if isTesting {
		leadingFilepath = "../."
	}

	filepath := fmt.Sprintf("%s./test_data/account_test_data.txt", leadingFilepath)
	account, err := accountProvider.ReadAccountFromFile(filepath)

	if err != nil {
		return nil, fmt.Errorf("error reading from test data file: %s", err)
	}

	return account, nil
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
