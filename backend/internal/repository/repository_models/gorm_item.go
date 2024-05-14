package repositorymodels

import (
	"github.com/lib/pq"
)

type GormItem struct {
	ID           int `gorm:"primaryKey"`
	ChatLink     string
	Name         string
	Icon         string
	Description  string
	Type         string
	Rarity       string
	Level        int
	VendorValue  int
	DefaultSkin  *int
	Flags        pq.StringArray  `gorm:"type:text[]"`
	GameTypes    pq.StringArray  `gorm:"type:text[]"`
	Restrictions pq.StringArray  `gorm:"type:text[]"`
	UpgradesInto *pq.StringArray `gorm:"type:text[]"`
	UpgradesFrom *pq.StringArray `gorm:"type:text[]"`
	// Details Details
}

// type GameType string

// type Flag string

// const (
// 	AccountBindOnUse
//     AccountBound
//     Attuned
//     BulkConsume
//     DeleteWarning
//     HideSuffix
//     Infused
//     MonsterOnly
//     NoMysticForge
//     NoSalvage
//     NoSell
//     NotUpgradeable
//     NoUnderwater
//     SoulbindOnAcquire
//     SoulBindOnUse
//     Tonic
//     Unique
// )

// type Details interface {

// }

// type ArmorDetails struct {
// 	Type: string
// 	WeightClass: string
// 	Defense: int
// 	InfusionSlots: InfusionSlot
// 	AttributeAdjustment: int
// 	InfixUpgrade: InfixUpgrade
// 	SuffixItemId: int
// 	SecondarySuffixItemId: string
// 	StatChoices: []int
// }

// func GormItemToItem(gormItem GormItem) models.Item {
// 	return models.Item{
// 		ID:           gormItem.ID,
// 		ChatLink:     gormItem.Icon,
// 		Name:         gormItem.Icon,
// 		Icon:         gormItem.Icon,
// 		Description:  gormItem.Icon,
// 		Type:         gormItem.Icon,
// 		Rarity:       gormItem.Icon,
// 		Level:        gormItem.Level,
// 		VendorValue:  gormItem.VendorValue,
// 		DefaultSkin:  *gormItem.DefaultSkin,
// 		Flags:        gormItem.Flags,
// 		GameTypes:    gormItem.GameTypes,
// 		Restrictions: gormItem.Restrictions,
// 		UpgradesInto: *gormItem.UpgradesInto,
// 		UpgradesFrom: *gormItem.UpgradesFrom,
// 		// Details : gormItem.Details,
// 	}
// }

// func ItemToGormItem(item models.Item) GormItem {
// 	return GormItem{
// 		ID:           item.ID,
// 		ChatLink:     item.Icon,
// 		Name:         item.Icon,
// 		Icon:         item.Icon,
// 		Description:  item.Icon,
// 		Type:         item.Icon,
// 		Rarity:       item.Icon,
// 		Level:        item.Level,
// 		VendorValue:  item.VendorValue,
// 		DefaultSkin:  &item.DefaultSkin,
// 		Flags:        item.Flags,
// 		GameTypes:    item.GameTypes,
// 		Restrictions: item.Restrictions,
// 		UpgradesInto: &item.UpgradesInto,
// 		UpgradesFrom: &item.UpgradesFrom,
// 		// Details : item.Details,
// 	}
// }
