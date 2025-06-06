package gw2models

import (
	"github.com/lib/pq"
	dbmodels "github.com/zoehay/gw2armoury/backend/internal/db/models"
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
	Slot      *string                 `json:"slot"`
	Location  *string                 `json:"location"`
}

func (gw2BagItem GW2BagItem) ToDBBagItem(accountID string, apiCharacterName *string) dbmodels.DBBagItem {
	var stats = (*dbmodels.DetailsMap)(gw2BagItem.Stats)
	return dbmodels.DBBagItem{
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
		Slot:          gw2BagItem.Slot,
		Location:      gw2BagItem.Location,
	}

}
