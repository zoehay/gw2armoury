package gw2models

type GW2Character struct {
	Name string    `json:"name"`
	Bags *[]GW2Bag `json:"bags"`
}
