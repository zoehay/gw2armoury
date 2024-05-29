package apimodels

import (
	"github.com/lib/pq"
	repositorymodels "github.com/zoehay/gw2armoury/backend/internal/repository/repository_models"
)

type ApiBag struct {
	Id        int           `json:"id"`
	Size      int           `json:"size"`
	Inventory []*ApiBagItem `json:"inventory"`
}

type ApiBagItem struct {
	Id        int                     `json:"id"`
	Count     int                     `json:"count"`
	Charges   *int                    `json:"charges,omitempty"`
	Infusions *[]int64                `json:"infusions,omitempty"`
	Upgrades  *[]int64                `json:"upgrades,omitempty"`
	Skin      *int                    `json:"skin,omitempty"`
	Stats     *map[string]interface{} `json:"stats,omitempty"`
	Dyes      *[]int64                `json:"dyes,omitempty"`
	Binding   *string                 `json:"binding,omitempty"`
	BoundTo   *string                 `json:"bound_to,omitempty"`
}

func ApiBagToGormBagItem(apiCharacterName string, apiBagItem ApiBagItem) repositorymodels.GormBagItem {
	var stats = (*repositorymodels.DetailsMap)(apiBagItem.Stats)
	return repositorymodels.GormBagItem{
		CharacterName: apiCharacterName,
		BagItemId:     apiBagItem.Id,
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
