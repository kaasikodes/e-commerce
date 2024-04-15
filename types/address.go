package types

import "github.com/kaasikodes/e-commerce-go/models"

type AddressInput struct {
	StreetAddress string `json:"streetAddress" validate:"required"`
	LgaID         string `json:"lgaId" validate:"required"`
	StateID       string `json:"stateId" validate:"required"`
	CountryID     string `json:"countryId" validate:"required"`
}
type PaginatedAddressesDataOutput struct {
	Data       []models.Address `json:"data"`
	NextCursor string   `json:"nextCursor"`
	HasMore    bool     `json:"hasMore"`
	Total      int      `json:"total"`
}
type RetrievAddressesInput struct {
	Pagination Pagination
}

type AddressRepository interface {
	CreateAddress(data AddressInput) (addressId string, errror error)
	RetrieveAddresses(input RetrievAddressesInput) (PaginatedAddressesDataOutput, error)
}