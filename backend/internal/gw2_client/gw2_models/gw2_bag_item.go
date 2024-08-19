package gw2models

import (
	"github.com/lib/pq"
	repositorymodels "github.com/zoehay/gw2armoury/backend/internal/repository/repository_models"
)

type GW2BagItem struct {
	ID        uint                    `json:"id"`
	Count     uint                    `json:"count"`
	Charges   *uint                   `json:"charges,omitempty"`
	Infusions *[]int64                `json:"infusions,omitempty"`
	Upgrades  *[]int64                `json:"upgrades,omitempty"`
	Skin      *uint                   `json:"skin,omitempty"`
	Stats     *map[string]interface{} `json:"stats,omitempty"`
	Dyes      *[]int64                `json:"dyes,omitempty"`
	Binding   *string                 `json:"binding,omitempty"`
	BoundTo   *string                 `json:"bound_to,omitempty"`
}

func (gw2BagItem GW2BagItem) ToGORMBagItem(accountID string, apiCharacterName string) repositorymodels.GORMBagItem {
	var stats = (*repositorymodels.DetailsMap)(gw2BagItem.Stats)
	return repositorymodels.GORMBagItem{
		AccountID:     accountID,
		CharacterName: apiCharacterName,
		BagItemID:     gw2BagItem.ID,
		Count:         gw2BagItem.Count,
		Charges:       gw2BagItem.Charges,
		Infusions:     (*pq.Int64Array)(gw2BagItem.Infusions),
		Upgrades:      (*pq.Int64Array)(gw2BagItem.Infusions),
		Skin:          gw2BagItem.Skin,
		Stats:         stats,
		Dyes:          (*pq.Int64Array)(gw2BagItem.Infusions),
		Binding:       gw2BagItem.Binding,
		BoundTo:       gw2BagItem.BoundTo,
	}

}

// func ApiInfusionsToGormInfusions