package gw2models

import "github.com/zoehay/gw2armoury/backend/internal/api/models"

type GW2Account struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}

func (gw2Account GW2Account) ToAccount() models.Account {
	return models.Account{
		AccountID:   *gw2Account.ID,
		AccountName: gw2Account.Name,
	}
}
