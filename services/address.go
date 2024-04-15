package services

import (
	"context"
	"database/sql"

	"github.com/kaasikodes/e-commerce-go/constants"
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

func (r *AddressRepository) CreateAddress(data types.AddressInput, ) (addressId string, errror error) {
	db := r.db
	// prepare query
	query := `INSERT INTO Address (ID, StreetAddress, LgaID, StateID, CountryID) VALUES (?, ?, ?, ?, ?)`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return "", err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	addressId, _ = utils.GenerateRandomID(10)
	res, err := stmt.ExecContext(ctx, addressId, data.StreetAddress, data.LgaID, data.StateID, data.CountryID)
	if err !=nil {
		return "", err
	}
	_, err = res.RowsAffected()
	if err !=nil {
		return "", err
	}
	return addressId, nil

}