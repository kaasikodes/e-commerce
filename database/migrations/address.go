package migrations

import (
	"context"
	"database/sql"
	"time"

	"github.com/kaasikodes/e-commerce-go/utils"
)

func CreateCountryTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS Country (
		ID VARCHAR(255) PRIMARY KEY,
		Name VARCHAR(255) NOT NULL
	)`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, query)
	return utils.ErrHandler(err)

}
func CreateStateTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS State (
		ID VARCHAR(255) PRIMARY KEY,
		Name VARCHAR(255) NOT NULL,
		CountryID VARCHAR(255) NOT NULL,
		FOREIGN KEY (CountryID) REFERENCES Country(ID)
	)`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, query)
	return utils.ErrHandler(err)

}
func CreateLgaTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS Lga (
		ID VARCHAR(255) PRIMARY KEY,
		Name VARCHAR(255) NOT NULL,
		StateID VARCHAR(255) NOT NULL,
		FOREIGN KEY (StateID) REFERENCES State(ID)
	)`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, query)
	return utils.ErrHandler(err)

}
func CreateAddressTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS Address (
		ID VARCHAR(255) PRIMARY KEY,
		StreetAddress VARCHAR(255) NOT NULL,
		LgaID VARCHAR(255) NOT NULL,
		StateID VARCHAR(255) NOT NULL,
		CountryID VARCHAR(255) NOT NULL,
		FOREIGN KEY (LgaID) REFERENCES Lga(ID),
		FOREIGN KEY (StateID) REFERENCES State(ID),
		FOREIGN KEY (CountryID) REFERENCES Country(ID)
	)`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, query)
	return utils.ErrHandler(err)

}