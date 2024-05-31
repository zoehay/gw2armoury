package apimodels

type APICharacter struct {
	Name string    `json:"name"`
	Bags *[]APIBag `json:"bags"`
}
