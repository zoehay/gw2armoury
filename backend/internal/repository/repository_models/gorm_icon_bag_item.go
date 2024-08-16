package repositorymodels

import "github.com/zoehay/gw2armoury/backend/internal/models"

type GORMIconBagItem struct {
	AccountID     string
	CharacterName string
	BagItemID     uint                    `json:"id"`
	Icon          string                  `json:"icon"`
	Count         uint                    `json:"count"`
	Charges       *uint                   `json:"charges,omitempty"`
	Infusions     *[]int64                `json:"infusions,omitempty" gorm:"type:integer[]"`
	Upgrades      *[]int64                `json:"upgrades,omitempty" gorm:"type:integer[]"`
	Skin          *uint                   `json:"skin,omitempty"`
	Stats         *map[string]interface{} `json:"stats,omitempty" gorm:"type:json"`
	Dyes          *[]int64                `json:"dyes,omitempty" gorm:"type:integer[]"`
	Binding       *string                 `json:"binding,omitempty"`
	BoundTo       *string                 `json:"bound_to,omitempty"`
}

func (gormIconBagItem GORMIconBagItem) ToBagItem() models.BagItem {
	return models.BagItem{
		// AccountID:     gormIconBagItem.AccountID,
		// CharacterName: gormIconBagItem.CharacterName,
		BagItemID: gormIconBagItem.BagItemID,
		Icon:      gormIconBagItem.Icon,
		Count:     gormIconBagItem.Count,
		Charges:   gormIconBagItem.Charges,
		Infusions: gormIconBagItem.Infusions,
		Upgrades:  gormIconBagItem.Upgrades,
		Skin:      gormIconBagItem.Skin,
		Stats:     gormIconBagItem.Stats,
		Dyes:      gormIconBagItem.Dyes,
		Binding:   gormIconBagItem.Binding,
		BoundTo:   gormIconBagItem.BoundTo,
	}
}
