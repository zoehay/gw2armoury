package apimodels

import (
	"github.com/zoehay/gw2armoury/backend/internal/models"
	repositorymodels "github.com/zoehay/gw2armoury/backend/internal/repository/repository_models"
)

type ItemRequest struct {
	Ids string `json:"ids"`
}

type ItemResponse struct {
	Items []ApiItem
}

type ApiItem struct {
	ID           int       `json:"id"`
	ChatLink     string    `json:"chat_link"`
	Name         string    `json:"name"`
	Icon         string    `json:"icon"`
	Description  string    `json:"description"`
	Type         string    `json:"type"`
	Rarity       string    `json:"rarity"`
	Level        int       `json:"level"`
	VendorValue  int       `json:"vendor_value"`
	DefaultSkin  *int      `json:"default_skin,omitempty"`
	Flags        []string  `json:"flags"`
	GameTypes    []string  `json:"game_types"`
	Restrictions []string  `json:"restrictions"`
	UpgradesInto *[]string `json:"upgrades_into,omitempty"`
	UpgradesFrom *[]string `json:"upgrades_from,omitempty"`
	// Details map[string]string `json:"details,omitempty"`
}

//   type Gw2ItemDetails struct {

//   }

func ApiItemToItem(apiItem ApiItem) models.Item {
	return models.Item{
		ID:           apiItem.ID,
		ChatLink:     apiItem.ChatLink,
		Name:         apiItem.Name,
		Icon:         apiItem.Icon,
		Description:  apiItem.Description,
		Type:         apiItem.Type,
		Rarity:       apiItem.Rarity,
		Level:        apiItem.Level,
		VendorValue:  apiItem.VendorValue,
		DefaultSkin:  *apiItem.DefaultSkin,
		Flags:        apiItem.Flags,
		GameTypes:    apiItem.GameTypes,
		Restrictions: apiItem.Restrictions,
		UpgradesInto: *apiItem.UpgradesInto,
		UpgradesFrom: *apiItem.UpgradesFrom,
		// Details : apiItem.Details,
	}
}

func ApiItemToGormItem(apiItem ApiItem) repositorymodels.GormItem {
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
		Flags:        apiItem.Flags,
		GameTypes:    apiItem.GameTypes,
		Restrictions: apiItem.Restrictions,
		UpgradesInto: apiItem.UpgradesInto,
		UpgradesFrom: apiItem.UpgradesFrom,
		// Details : apiItem.Details,
	}
}

type ItemError struct {
	Text string `json:"text"`
}
