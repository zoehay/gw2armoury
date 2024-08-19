package gw2models

type GW2Bag struct {
	ID        uint          `json:"id"`
	Size      uint          `json:"size"`
	Inventory []*GW2BagItem `json:"inventory"`
}
