package migrations

import (
	"context"
	"database/sql"
	"time"

	"github.com/kaasikodes/e-commerce-go/utils"
)


func CreateCartTable (db *sql.DB) error{
	query := `CREATE TABLE IF NOT EXISTS Cart (
		ID VARCHAR(255) PRIMARY KEY,
		CustomerID VARCHAR(255) NOT NULL,
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (CustomerID) REFERENCES Customer(ID)
	)`

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx,query)
	return utils.ErrHandler(err)

}
func CreateCartItemTable (db *sql.DB) error{
	query := `CREATE TABLE IF NOT EXISTS CartItem (
		ID VARCHAR(255) PRIMARY KEY,
		ProductID VARCHAR(255) NOT NULL,
		CartID VARCHAR(255) NOT NULL,
		Quantity INT NOT NULL,
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (ProductID) REFERENCES Product(ID),
		FOREIGN KEY (CartID) REFERENCES Cart(ID)
	)`

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx,query)
	return utils.ErrHandler(err)

}