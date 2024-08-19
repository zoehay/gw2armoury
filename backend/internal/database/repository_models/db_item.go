package repositorymodels

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"github.com/zoehay/gw2armoury/backend/internal/api/models"
)

type DBItem struct {
	Name         string
	Description  string
	Type         string
	Level        uint
	Rarity       string
	VendorValue  uint
	DefaultSkin  *uint
	GameTypes    pq.StringArray `gorm:"type:text[]"`
	Flags        pq.StringArray `gorm:"type:text[]"`
	Restrictions pq.StringArray `gorm:"type:text[]"`
	ID           uint           `gorm:"primaryKey"`
	ChatLink     string
	Icon         string
	UpgradesInto *pq.StringArray `gorm:"type:text[]"`
	UpgradesFrom *pq.StringArray `gorm:"type:text[]"`
	Details      *DetailsMap     `gorm:"type:json"`
}

func (dbItem DBItem) ToItem() models.Item {
	return models.Item{
		Name:         dbItem.Name,
		Type:         dbItem.Name,
		Level:        dbItem.Level,
		Rarity:       dbItem.Rarity,
		VendorValue:  dbItem.VendorValue,
		DefaultSkin:  dbItem.DefaultSkin,
		GameTypes:    dbItem.GameTypes,
		Flags:        dbItem.Flags,
		Restrictions: dbItem.Restrictions,
		ID:           dbItem.ID,
		ChatLink:     dbItem.ChatLink,
		Icon:         dbItem.Icon,
		Description:  dbItem.Description,
		UpgradesInto: (*[]string)(dbItem.UpgradesInto),
		UpgradesFrom: (*[]string)(dbItem.UpgradesFrom),
		Details:      (*map[string]interface{})(dbItem.Details),
	}
}

type DetailsMap map[string]interface{}

func (detailsMap *DetailsMap) Scan(value interface{}) error {
	if value == nil {
		*detailsMap = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	err := json.Unmarshal(bytes, detailsMap)
	return err

}

func (detailsMap DetailsMap) Value() (driver.Value, error) {
	if len(detailsMap) == 0 {
		return nil, nil
	}
	return json.Marshal(detailsMap)
}

// type GormArmorDetails struct {
// 	Type        string
// 	WeightClass string
// 	Defense     int
// 	// InfusionSlots       *[]InfusionSlotType `json:"infusion_slots"`
// 	AttributeAdjustment float32 `sql:"type:decimal(10,2);"`
// 	// InfixUpgrade        *ApiInfixUpgrade `json:"infix_upgrade"`
// 	SuffixItemId          *int
// 	SecondarySuffixItemId string
// 	// StatChoices *[]int `json:"stat_choices"`
// }

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
