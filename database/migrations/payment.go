package migrations

import (
	"context"
	"database/sql"
	"time"

	"github.com/kaasikodes/e-commerce-go/utils"
)

func CreatePaymentTable (db *sql.DB) error{
	query := `CREATE TABLE IF NOT EXISTS Payment (
		ID VARCHAR(255) PRIMARY KEY,
		OrderID VARCHAR(255) NOT NULL,
		Amount INT NOT NULL,
		Paid BOOLEAN NOT NULL,
		PaidAt TIMESTAMP,
		Method VARCHAR(255) NOT NULL,
		FOREIGN KEY (OrderID) REFERENCES ` + "`Order`" + `(ID)
	)`

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx,query)
	return utils.ErrHandler(err)

}
