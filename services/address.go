package services

import (
	"database/sql"

	"github.com/kaasikodes/e-commerce-go/types"
	"github.com/kaasikodes/e-commerce-go/utils"
)

type AddressRepository struct {
	db *sql.DB
}

func NewAddressRepository(db *sql.DB) *AddressRepository {
	return &AddressRepository{
		db: db,
	}
}

func (r *AddressRepository) CreateAddress(data types.CreateOrderInput, ) (addressId string, errror error) {
	addressId, err := utils.GenerateRandomID(10)
	if err != nil {
		return "", err
	}
	return addressId, nil
}