package migrations

import (
	"context"
	"database/sql"

	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/utils"
)

func CreatePasswordResetTokenTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS  PasswordResetToken (
		ID VARCHAR(255) PRIMARY KEY,
		Token VARCHAR(255) NOT NULL,
		Email VARCHAR(255) NOT NULL UNIQUE,
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		ExpiresAt TIMESTAMP,
		FOREIGN KEY (Email) REFERENCES User(Email)
	)
	`

	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()

	_, err := db.ExecContext(ctx, query)
	return utils.ErrHandler(err)

}
func CreateVerificationTokenTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS  VerificationToken (
		ID VARCHAR(255) PRIMARY KEY,
		Token VARCHAR(255) NOT NULL,
		Email VARCHAR(255) NOT NULL UNIQUE,
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		ExpiresAt TIMESTAMP,
		FOREIGN KEY (Email) REFERENCES User(Email)
	)
	`

	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()

	_, err := db.ExecContext(ctx, query)
	return utils.ErrHandler(err)

}