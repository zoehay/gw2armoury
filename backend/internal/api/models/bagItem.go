package models

type BagItem struct {
	CharacterName string                  `json:"character_name"`
	Name          *string                 `json:"name"`
	Description   *string                 `json:"description"`
	BagItemID     uint                    `json:"id"`
	Icon          string                  `json:"icon"`
	Count         uint                    `json:"count"`
	Charges       *uint                   `json:"charges,omitempty"`
	Infusions     *[]int64                `json:"infusions,omitempty"`
	Upgrades      *[]int64                `json:"upgrades,omitempty"`
	Skin          *uint                   `json:"skin,omitempty"`
	Stats         *map[string]interface{} `json:"stats,omitempty" gorm:"type:json"`
	Dyes          *[]int64                `json:"dyes,omitempty" gorm:"type:integer[]"`
	Binding       *string                 `json:"binding,omitempty"`
	BoundTo       *string                 `json:"bound_to,omitempty"`
}
