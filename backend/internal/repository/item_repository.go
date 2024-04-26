package repository

import (
	"github.com/zoehay/gw2armoury/backend/internal/models"
)

type ItemRepository struct {
	GormRepository
}

type gormItem struct {
	ID int `gorm:"primaryKey"`
	ChatLink string
	Name string
	Icon string
	Description string
	Type string
	Rarity string
	Level uint
	VendorValue uint
	DefaultSkin uint
	Flags []string `gorm:"type:text"`
	GameTypes []string `gorm:"type:text"`
	Restrictions []string `gorm:"type:text"`
	UpgradesInto []string `gorm:"type:text"`
	UpgradesFrom []string `gorm:"type:text"`
	Details string
  }


func (repository *ItemRepository) Create(item models.Item) (*models.Item, error) {
	gormItem := gormItem{
		ID: item.ID,
		ChatLink: item.ChatLink, 
		Name: item.Name,
		Icon: item.Icon,
		Description: item.Description,
		Type: item.Type,
		Rarity: item.Rarity,
		Level: item.Level,
		VendorValue: item.VendorValue,
		DefaultSkin: item.DefaultSkin,
		Flags: item.Flags,
		GameTypes: item.GameTypes,
		Restrictions: item.Restrictions,
		UpgradesInto: item.UpgradesInto,
		UpgradesFrom: item.UpgradesFrom,
		Details: item.Details,
	}

	err := repository.db.Create(&gormItem).Error
	if err != nil {
		return nil, err
	}

	result := models.Item(gormItem)
	return &result, nil
	
}

func (repository *ItemRepository) GetAll() ([]models.Item, error) {
	var gormItems []models.Item

	err := repository.db.Find(&gormItems).Error
	if err != nil {
		return nil, err
	}

	var result []models.Item
	for _, gormItem := range gormItems {
		result = append(result, models.Item(gormItem))
	}

	return result, nil

}