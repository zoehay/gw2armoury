package gw2apimodels

type ItemRequest struct {
	Ids string `json:"ids"`
}

type ItemResponse struct {
	Items []Gw2Item
}

type Gw2Item struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Type string `json:"type"`
	Level uint `json:"level"`
	Rarity string `json:"rarity"`
	VendorValue uint `json:"vendor_value"`
	GameTypes []string `json:"game_types"`
	Flags []string `json:"flags"`
	Restrictions []string `json:"restrictions"`
	ID int `json:"id"`
	ChatLink string `json:"chat_link"`
	Icon string `json:"icon"`
	DefaultSkin *int `json:"default_skin,omitempty"`
	UpgradesInto *[]string `json:"upgrades_into,omitempty"`
	UpgradesFrom *[]string `json:"upgrades_from,omitempty"`
	// Details map[string]string `json:"details,omitempty"`
  }

//   type Gw2ItemDetails struct {

//   }

  type ItemError struct {
	Text string `json:"text"`
  }