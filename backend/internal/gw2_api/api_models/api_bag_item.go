package apimodels

import (
	"github.com/lib/pq"
	repositorymodels "github.com/zoehay/gw2armoury/backend/internal/repository/repository_models"
)

type APIBagItem struct {
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

func APIBagToGORMBagItem(accountID string, apiCharacterName string, apiBagItem APIBagItem) repositorymodels.GORMBagItem {
	var stats = (*repositorymodels.DetailsMap)(apiBagItem.Stats)
	return repositorymodels.GORMBagItem{
		AccountID:     accountID,
		CharacterName: apiCharacterName,
		BagItemID:     apiBagItem.ID,
		Count:         apiBagItem.Count,
		Charges:       apiBagItem.Charges,
		Infusions:     (*pq.Int64Array)(apiBagItem.Infusions),
		Upgrades:      (*pq.Int64Array)(apiBagItem.Infusions),
		Skin:          apiBagItem.Skin,
		Stats:         stats,
		Dyes:          (*pq.Int64Array)(apiBagItem.Infusions),
		Binding:       apiBagItem.Binding,
		BoundTo:       apiBagItem.BoundTo,
	}

}

// func ApiInfusionsToGormInfusions
