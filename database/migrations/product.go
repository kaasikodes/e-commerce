package migrations

import (
	"context"
	"database/sql"
	"time"

	"github.com/kaasikodes/e-commerce-go/utils"
)


func CreateProductTable (db *sql.DB) error{
	query := `CREATE TABLE IF NOT EXISTS Product  (
		ID VARCHAR(255) PRIMARY KEY,
		Name VARCHAR(255) NOT NULL,
		Description TEXT,
		Price INT NOT NULL,
		Quantity INT NOT NULL,
		CategoryID VARCHAR(255) NOT NULL,
		OwnerID VARCHAR(255) NOT NULL,
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (CategoryID) REFERENCES Category(ID),
		FOREIGN KEY (OwnerID) REFERENCES Seller(ID)
	)`

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx,query)
	return utils.ErrHandler(err)

}