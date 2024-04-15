package seeders

import (
	"context"
	"database/sql"
	"strings"

	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/utils"
)

// add values to country
func AddCountries(db *sql.DB) error {
	data := []struct {
		ID   string
		Name string
	}{
		{"14", "Country1"},
		{"5", "Country2"},
		{"8", "Country3"},
	}

	query := `INSERT INTO Country (ID, Name) VALUES`

	var inserts []string
	var params []interface{}
	for _, country := range data {
		inserts = append(inserts, "(?, ?)")
		params = append(params, country.ID, country.Name)
	}
	queryVals := strings.Join(inserts, ",")
	query = query + queryVals

	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, params...)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// add values to state
func AddStates(db *sql.DB) error {
	data := []struct {
		ID string
		Name      string
		CountryID string
	}{
		{"14", "State1", "14"},
		{"5", "State2", "5"},
		{"8", "State3", "8"},
	}

	query := `INSERT INTO State (ID, Name, CountryID) VALUES`

	var inserts []string
	var params []interface{}
	for _, state := range data {
		inserts = append(inserts, "(?, ?, ?)")
		params = append(params, state.ID, state.Name, state.CountryID)
	}
	queryVals := strings.Join(inserts, ",")
	query = query + queryVals

	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, params...)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// add values to lga
func AddLGAs(db *sql.DB) error {
	data := []struct {
		ID string
		Name      string
		StateID string
	}{
		{"14", "Lga1", "14"},
		{"5", "Lga2", "5"},
		{"8", "Lga3", "8"},
	}

	query := `INSERT INTO Lga (ID, Name, StateID) VALUES`
	var inserts []string
	var params []interface{}
	for _, lga := range data {
		inserts = append(inserts, "(?, ?, ?)")
		
		params = append(params, lga.ID, lga.Name, lga.StateID) 
	}
	queryVals := strings.Join(inserts, ",")
	query = query + queryVals

	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, params...)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// add values to address
func AddAddresses(db *sql.DB) error {
	data := []struct {
		StreetAddress   string
		LgaID     string
		StateID  string
		CountryID string
	}{
		{"Street1", "14", "14", "14"},
		{"Street2", "5", "5", "5"},
		{"Street3", "8", "8", "8"},
	}

	query := `INSERT INTO Address (ID, StreetAddress, LgaID, StateID, CountryID) VALUES`

	var inserts []string
	var params []interface{}
	for _, address := range data {
		inserts = append(inserts, "(?, ?, ?, ?, ?)")
		id, _ := utils.GenerateRandomID(10)
		params = append(params, id, address.StreetAddress, address.LgaID, address.StateID, address.CountryID)
	}
	queryVals := strings.Join(inserts, ",")
	query = query + queryVals

	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, params...)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
