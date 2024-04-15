package services

import (
	"context"
	"database/sql"

	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/models"
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

func (r *AddressRepository) RetrieveAddresses(input types.RetrievAddressesInput) (types.PaginatedAddressesDataOutput, error) {
	db := r.db
	// prepare query - get  address, join stateID, lgaID
	query := `
	SELECT a.ID, a.StreetAddress, a.LgaID, a.StateID, a.CountryID, l.ID, l.Name, s.ID, s.Name, c.ID, c.Name,
	(SELECT COUNT(*) FROM Address) AS total_addresses
	FROM Address a
	JOIN Lga l ON l.ID = a.LgaID
	JOIN State s ON s.ID = l.StateID
	JOIN Country c ON c.ID = s.CountryID
	WHERE a.ID > ? 
    ORDER BY a.ID ASC
    LIMIT ?

	`


	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	output := types.PaginatedAddressesDataOutput{}
	stmt, err := db.PrepareContext(ctx, query )
	if err !=nil {
		return output, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	addresses := []models.Address{}
	rows, err := stmt.QueryContext(ctx, utils.Ternary(input.Pagination.NextCursor == "", "", input.Pagination.NextCursor), utils.Ternary(input.Pagination.PageSize == 0, constants.DefaultPageSize, input.Pagination.PageSize))
	if err !=nil {
		return output, err
	}
	defer rows.Close()
	total := 0
	for rows.Next() {
		address := models.Address{}
		address.Country = &models.Country{}
		address.State = &models.State{}
		address.Lga = &models.Lga{}
		err = rows.Scan(&address.ID, &address.StreetAddress, &address.LgaID, &address.StateID, &address.CountryID, &address.Lga.ID, &address.Lga.Name, &address.State.ID, &address.State.Name, &address.Country.ID, &address.Country.Name, &total)
		if err !=nil {
			return output, err
		}
		addresses = append(addresses, address)
	}
	lastItemId := ""
	// select last item in the list
	if len(addresses) > 0 {
		lastItem := addresses[len(addresses)-1]
		lastItemId = lastItem.ID
	}

	output = types.PaginatedAddressesDataOutput{
		Data: addresses,
		NextCursor: lastItemId,
		HasMore:    len(addresses) < total,
		Total:      total,
	}
	return output, nil

}
