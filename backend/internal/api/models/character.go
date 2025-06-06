package models

type Character struct {
	Name      string    `json:"name"`
	Equipment []BagItem `json:"equipment"`
	Inventory []BagItem `json:"inventory"`
}
