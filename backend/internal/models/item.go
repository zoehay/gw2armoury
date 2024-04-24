package models

type Item struct {
  ID uint
  ChatLink string
  Name string
  Icon string
  Description string
  Type string
  Rarity string
  Level uint
  VendorValue uint
  DefaultSkin uint
  Flags []string
  GameTypes []string
  Restrictions []string
  UpgradesInto []string
  UpgradesFrom []string
  Details string;
}