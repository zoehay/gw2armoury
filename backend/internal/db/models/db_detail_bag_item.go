package dbmodels

import "github.com/zoehay/gw2armoury/backend/internal/api/models"

type DBDetailBagItem struct {
	AccountID     string                  `json:"account_id"`
	CharacterName string                  `json:"character_name"`
	Name          *string                 `json:"name"`
	Description   *string                 `json:"description"`
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
	Rarity        *string                 `json:"rarity"`
	Slot          *string                 `json:"slot"`
	Location      *string                 `json:"location"`
}

func (dbDetailBagItem DBDetailBagItem) ToBagItem() models.BagItem {
	return models.BagItem{
		CharacterName: dbDetailBagItem.CharacterName,
		Name:          dbDetailBagItem.Name,
		Description:   dbDetailBagItem.Description,
		BagItemID:     dbDetailBagItem.BagItemID,
		Icon:          dbDetailBagItem.Icon,
		Count:         dbDetailBagItem.Count,
		Charges:       dbDetailBagItem.Charges,
		Infusions:     dbDetailBagItem.Infusions,
		Upgrades:      dbDetailBagItem.Upgrades,
		Skin:          dbDetailBagItem.Skin,
		Stats:         dbDetailBagItem.Stats,
		Dyes:          dbDetailBagItem.Dyes,
		Binding:       dbDetailBagItem.Binding,
		BoundTo:       dbDetailBagItem.BoundTo,
		Rarity:        dbDetailBagItem.Rarity,
		Slot:          dbDetailBagItem.Slot,
		Location:      dbDetailBagItem.Location,
	}
}

func DBDetailBagItemsToAccountInventory(dbIconBagItems []DBDetailBagItem, accountID string) models.AccountInventory {

	characterNameMap := map[string]models.Character{}
	var sharedInventory []models.BagItem
	var characters []models.Character

	for i := range dbIconBagItems {
		item := dbIconBagItems[i].ToBagItem()
		name := item.CharacterName

		if name == "Shared Inventory" {
			sharedInventory = append(sharedInventory, item)
		} else {
			entry, ok := characterNameMap[name]
			isEquipment := item.IsEquipment()

			if ok {
				if isEquipment {
					entry.Equipment = append(entry.Equipment, item)
					characterNameMap[name] = entry
				} else {
					entry.Inventory = append(entry.Inventory, item)
					characterNameMap[name] = entry
				}
			} else {
				newCharacter := &models.Character{
					Name:      name,
					Equipment: []models.BagItem{},
					Inventory: []models.BagItem{},
				}
				if isEquipment {
					newCharacter.Equipment = append(newCharacter.Equipment, item)
				} else {
					newCharacter.Inventory = append(newCharacter.Inventory, item)
				}
				characterNameMap[name] = *newCharacter
			}
		}

	}

	for character := range characterNameMap {
		characters = append(characters, characterNameMap[character])
	}

	var accountInventory models.AccountInventory
	accountInventory.AccountID = accountID
	accountInventory.SharedInventory = &sharedInventory
	accountInventory.Characters = &characters

	return accountInventory
}
