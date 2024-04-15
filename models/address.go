package models

type Address struct {
	ID            string `json:"id"`
	StreetAddress string `json:"streetAddress"`
	LgaID         string `json:"lgaId"`
	Lga           *Lga
	State         *State
	StateID       string `json:"stateId"`
	CountryID     string `json:"countryId"`
	Country       *Country
}

type Lga struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	State   State
	StateId string `json:"stateId"`
}
type State struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Lgas      []Lga  `json:"lgas"`
	CountryID string `json:"countryId"`
}
type Country struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	States []State `json:"states"`
}