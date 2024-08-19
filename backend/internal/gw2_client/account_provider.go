package gw2api

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/zoehay/gw2armoury/backend/internal/clients"
	gw2models "github.com/zoehay/gw2armoury/backend/internal/gw2_client/gw2_models"
)

type AccountDataProvider interface {
	GetAccount(apiKey string) (*gw2models.GW2Account, error)
}

type AccountProvider struct{}

func (accountProvider *AccountProvider) GetAccount(apiKey string) (*gw2models.GW2Account, error) {
	res, err := clients.GetAccountID(apiKey)

	if err != nil {
		return nil, fmt.Errorf("provider get error: %s", err)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("provider io.ReadAll error: %s", err)
	}

	var result gw2models.GW2Account

	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("provider json.Unmarshal error: %s", err)
	}

	return &result, nil

}
