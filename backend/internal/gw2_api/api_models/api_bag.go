package apimodels

type APIBag struct {
	ID        uint          `json:"id"`
	Size      uint          `json:"size"`
	Inventory []*APIBagItem `json:"inventory"`
}
