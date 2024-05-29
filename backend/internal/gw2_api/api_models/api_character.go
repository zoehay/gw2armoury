package apimodels

type ApiCharacter struct {
	Name string    `json:"name"`
	Bags *[]ApiBag `json:"bags"`
}
