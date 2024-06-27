package models

type Item struct {
	Name         string                  `json:"name"`
	Type         string                  `json:"type"`
	Level        uint                    `json:"level"`
	Rarity       string                  `json:"rarity"`
	VendorValue  uint                    `json:"vendor_value"`
	DefaultSkin  *uint                   `json:"default_skin,omitempty"`
	GameTypes    []string                `json:"game_types" gorm:"type:text[]"`
	Flags        []string                `json:"flags" gorm:"type:text[]"`
	Restrictions []string                `json:"restrictions" gorm:"type:text[]"`
	ID           uint                    `json:"id"`
	ChatLink     string                  `json:"chat_link"`
	Icon         string                  `json:"icon"`
	Description  string                  `json:"description"`
	UpgradesInto *[]string               `json:"upgrades_into,omitempty" gorm:"type:text[]"`
	UpgradesFrom *[]string               `json:"upgrades_from,omitempty" gorm:"type:text[]"`
	Details      *map[string]interface{} `json:"details,omitempty" gorm:"type:json"`
}
