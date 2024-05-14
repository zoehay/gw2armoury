package models

type Item struct {
	ID           int
	ChatLink     string
	Name         string
	Icon         string
	Description  string
	Type         string
	Rarity       string
	Level        int
	VendorValue  int
	DefaultSkin  int
	Flags        []string
	GameTypes    []string
	Restrictions []string
	UpgradesInto []string
	UpgradesFrom []string
	// Details map[string]string
}
