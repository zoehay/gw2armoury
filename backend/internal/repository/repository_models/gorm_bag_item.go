package repositorymodels

import "github.com/lib/pq"

// MVP do not care about which bag an item is in
type GormBagItem struct {
	CharacterName string
	BagItemId     uint
	Count         uint
	Charges       *uint
	Infusions     *pq.Int64Array `gorm:"type:integer[]"`
	Upgrades      *pq.Int64Array `gorm:"type:integer[]"`
	Skin          *uint
	Stats         *DetailsMap    `gorm:"type:json"`
	Dyes          *pq.Int64Array `gorm:"type:integer[]"`
	Binding       *string
	BoundTo       *string
}
