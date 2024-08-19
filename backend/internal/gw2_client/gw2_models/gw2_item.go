package gw2models

import (
	"github.com/lib/pq"
	repositorymodels "github.com/zoehay/gw2armoury/backend/internal/database/repository_models"
)

type ItemRequest struct {
	IDs string `json:"ids"`
}

type ItemResponse struct {
	Items []GW2Item
}

type GW2Item struct {
	Name         string                  `json:"name"`
	Type         string                  `json:"type"`
	Level        uint                    `json:"level"`
	Rarity       string                  `json:"rarity"`
	VendorValue  uint                    `json:"vendor_value"`
	DefaultSkin  *uint                   `json:"default_skin,omitempty"`
	GameTypes    []string                `json:"game_types"`
	Flags        []string                `json:"flags"`
	Restrictions []string                `json:"restrictions"`
	ID           uint                    `json:"id"`
	ChatLink     string                  `json:"chat_link"`
	Icon         string                  `json:"icon"`
	Description  string                  `json:"description"`
	UpgradesInto *[]string               `json:"upgrades_into,omitempty"`
	UpgradesFrom *[]string               `json:"upgrades_from,omitempty"`
	Details      *map[string]interface{} `json:"details,omitempty"`
}

func (gw2Item GW2Item) ToDBItem() repositorymodels.DBItem {
	var details = (*repositorymodels.DetailsMap)(gw2Item.Details)
	return repositorymodels.DBItem{
		ID:           gw2Item.ID,
		ChatLink:     gw2Item.ChatLink,
		Name:         gw2Item.Name,
		Icon:         gw2Item.Icon,
		Description:  gw2Item.Description,
		Type:         gw2Item.Type,
		Rarity:       gw2Item.Rarity,
		Level:        gw2Item.Level,
		VendorValue:  gw2Item.VendorValue,
		DefaultSkin:  gw2Item.DefaultSkin,
		Flags:        (pq.StringArray)(gw2Item.Flags),
		GameTypes:    (pq.StringArray)(gw2Item.GameTypes),
		Restrictions: (pq.StringArray)(gw2Item.Restrictions),
		UpgradesInto: (*pq.StringArray)(gw2Item.UpgradesInto),
		UpgradesFrom: (*pq.StringArray)(gw2Item.UpgradesFrom),
		Details:      details,
	}
}

// func (gw2Item GW2Item) ToItem(gw2Item gw2Item) models.Item {
// 	return models.Item{
// 		ID:           gw2Item.ID,
// 		ChatLink:     gw2Item.ChatLink,
// 		Name:         gw2Item.Name,
// 		Icon:         gw2Item.Icon,
// 		Description:  gw2Item.Description,
// 		Type:         gw2Item.Type,
// 		Rarity:       gw2Item.Rarity,
// 		Level:        gw2Item.Level,
// 		VendorValue:  gw2Item.VendorValue,
// 		DefaultSkin:  *gw2Item.DefaultSkin,
// 		Flags:        gw2Item.Flags,
// 		GameTypes:    gw2Item.GameTypes,
// 		Restrictions: gw2Item.Restrictions,
// 		UpgradesInto: *gw2Item.UpgradesInto,
// 		UpgradesFrom: *gw2Item.UpgradesFrom,
// 		Details:      gw2Item.Details,
// 	}
// }

type ItemError struct {
	Text string `json:"text"`
}
