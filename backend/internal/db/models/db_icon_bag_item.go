package dbmodels

import "github.com/zoehay/gw2armoury/backend/internal/api/models"

type DBIconBagItem struct {
	AccountID     string                  `json:"account_id"`
	CharacterName string                  `json:"character_name"`
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

func (dbIconBagItem DBIconBagItem) ToBagItem() models.BagItem {
	return models.BagItem{
		CharacterName: dbIconBagItem.CharacterName,
		BagItemID:     dbIconBagItem.BagItemID,
		Icon:          dbIconBagItem.Icon,
		Count:         dbIconBagItem.Count,
		Charges:       dbIconBagItem.Charges,
		Infusions:     dbIconBagItem.Infusions,
		Upgrades:      dbIconBagItem.Upgrades,
		Skin:          dbIconBagItem.Skin,
		Stats:         dbIconBagItem.Stats,
		Dyes:          dbIconBagItem.Dyes,
		Binding:       dbIconBagItem.Binding,
		BoundTo:       dbIconBagItem.BoundTo,
	}
}
