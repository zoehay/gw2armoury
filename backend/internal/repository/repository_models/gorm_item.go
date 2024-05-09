package repositorymodels

import (
	"github.com/zoehay/gw2armoury/backend/internal/models"
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
	DefaultSkin  int
	Flags        []string `gorm:"type:text"`
	GameTypes    []string `gorm:"type:text"`
	Restrictions []string `gorm:"type:text"`
	UpgradesInto []string `gorm:"type:text"`
	UpgradesFrom []string `gorm:"type:text"`
	// Details map[string]string
}

func GormItemToItem(gormItem GormItem) models.Item {
	return models.Item{
		ID:           gormItem.ID,
		ChatLink:     gormItem.Icon,
		Name:         gormItem.Icon,
		Icon:         gormItem.Icon,
		Description:  gormItem.Icon,
		Type:         gormItem.Icon,
		Rarity:       gormItem.Icon,
		Level:        gormItem.Level,
		VendorValue:  gormItem.VendorValue,
		DefaultSkin:  gormItem.DefaultSkin,
		Flags:        gormItem.Flags,
		GameTypes:    gormItem.GameTypes,
		Restrictions: gormItem.Restrictions,
		UpgradesInto: gormItem.UpgradesInto,
		UpgradesFrom: gormItem.UpgradesFrom,
		// Details : gormItem.Details,
	}
}

func ItemToGormItem(item models.Item) GormItem {
	return GormItem{
		ID:           item.ID,
		ChatLink:     item.Icon,
		Name:         item.Icon,
		Icon:         item.Icon,
		Description:  item.Icon,
		Type:         item.Icon,
		Rarity:       item.Icon,
		Level:        item.Level,
		VendorValue:  item.VendorValue,
		DefaultSkin:  item.DefaultSkin,
		Flags:        item.Flags,
		GameTypes:    item.GameTypes,
		Restrictions: item.Restrictions,
		UpgradesInto: item.UpgradesInto,
		UpgradesFrom: item.UpgradesFrom,
		// Details : item.Details,
	}
}
