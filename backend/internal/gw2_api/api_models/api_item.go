package apimodels

import (
	"github.com/lib/pq"
	repositorymodels "github.com/zoehay/gw2armoury/backend/internal/repository/repository_models"
)

type ItemRequest struct {
	Ids string `json:"ids"`
}

type ItemResponse struct {
	Items []ApiItem
}

type ApiItem struct {
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

// func ApiItemToItem(apiItem ApiItem) models.Item {
// 	return models.Item{
// 		ID:           apiItem.ID,
// 		ChatLink:     apiItem.ChatLink,
// 		Name:         apiItem.Name,
// 		Icon:         apiItem.Icon,
// 		Description:  apiItem.Description,
// 		Type:         apiItem.Type,
// 		Rarity:       apiItem.Rarity,
// 		Level:        apiItem.Level,
// 		VendorValue:  apiItem.VendorValue,
// 		DefaultSkin:  *apiItem.DefaultSkin,
// 		Flags:        apiItem.Flags,
// 		GameTypes:    apiItem.GameTypes,
// 		Restrictions: apiItem.Restrictions,
// 		UpgradesInto: *apiItem.UpgradesInto,
// 		UpgradesFrom: *apiItem.UpgradesFrom,
// 		Details:      apiItem.Details,
// 	}
// }

func ApiItemToGormItem(apiItem ApiItem) repositorymodels.GormItem {
	var details = (*repositorymodels.DetailsMap)(apiItem.Details)
	return repositorymodels.GormItem{
		ID:           apiItem.ID,
		ChatLink:     apiItem.ChatLink,
		Name:         apiItem.Name,
		Icon:         apiItem.Icon,
		Description:  apiItem.Description,
		Type:         apiItem.Type,
		Rarity:       apiItem.Rarity,
		Level:        apiItem.Level,
		VendorValue:  apiItem.VendorValue,
		DefaultSkin:  apiItem.DefaultSkin,
		Flags:        (pq.StringArray)(apiItem.Flags),
		GameTypes:    (pq.StringArray)(apiItem.GameTypes),
		Restrictions: (pq.StringArray)(apiItem.Restrictions),
		UpgradesInto: (*pq.StringArray)(apiItem.UpgradesInto),
		UpgradesFrom: (*pq.StringArray)(apiItem.UpgradesFrom),
		Details:      details,
	}
}

type ItemError struct {
	Text string `json:"text"`
}
