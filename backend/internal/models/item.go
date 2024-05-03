package models

type Item struct {
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