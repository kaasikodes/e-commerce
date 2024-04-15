package types

type AddressInput struct {
	StreetAddress string `json:"streetAddress" validate:"required"`
	LgaID         string `json:"lgaId" validate:"required"`
	StateID       string `json:"stateId" validate:"required"`
	CountryID     string `json:"countryId" validate:"required"`
}

type AddressRepository interface {
	CreateAddress(data AddressInput) (addressId string, errror error)
}