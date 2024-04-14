package types

import "github.com/kaasikodes/e-commerce-go/models"

type AddressInput struct {
	Street  string `json:"street" validate:"required"`
	LGA     string `json:"lga" validate:"required"`
	State   string `json:"state" validate:"required"`
	Country string `json:"country" validate:"required"`
}

type AddressRepository interface {
	CreateAddress(data CreateOrderInput) (addressId string, errror error)
	RetrieveAddress(id string) (models.Address, error)
}