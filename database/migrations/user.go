package migrations

import (
	"context"
	"database/sql"
	"time"

	"github.com/kaasikodes/e-commerce-go/utils"
)
func CreateUserTable (db *sql.DB) error{
	query := `CREATE TABLE IF NOT EXISTS  User (
		ID VARCHAR(255) PRIMARY KEY,
		Name VARCHAR(255) NOT NULL,
		Email VARCHAR(255) NOT NULL UNIQUE,
		Password VARCHAR(255) NOT NULL,
		RefreshToken VARCHAR(255),
		AccessToken VARCHAR(255),
		Roles VARCHAR(255) NOT NULL,
		Image VARCHAR(255),
		EmailVerified BOOLEAN,
		EmailVerifiedAt TIMESTAMP,
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	)
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx,query)
	return utils.ErrHandler(err)

}
func CreateSellerTable (db *sql.DB) error{
	query := `CREATE TABLE IF NOT EXISTS  Seller (
		ID VARCHAR(255) PRIMARY KEY,
		UserID VARCHAR(255) NOT NULL,
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (UserID) REFERENCES User(ID)
	)
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx,query)
	return utils.ErrHandler(err)

}
func CreateCustomerTable (db *sql.DB) error{
	query := `CREATE TABLE IF NOT EXISTS  Customer (
		ID VARCHAR(255) PRIMARY KEY,
		UserID VARCHAR(255) NOT NULL,
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (UserID) REFERENCES User(ID)
	)
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx,query)
	return utils.ErrHandler(err)

}