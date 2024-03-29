package migrations

import (
	"context"
	"database/sql"
	"time"

	"github.com/kaasikodes/e-commerce-go/utils"
)
func CreateOrderTable (db *sql.DB) error{
	query := "CREATE TABLE IF NOT EXISTS `Order` ( " +
	"ID VARCHAR(255) PRIMARY KEY, " +
	"UserID VARCHAR(255) NOT NULL, " +
	"TotalAmount INT NOT NULL, " +
	"DeliveryAddressID VARCHAR(255) NOT NULL, " +
	"CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, " +
	"UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, " +
	"FOREIGN KEY (UserID) REFERENCES User(ID), " +
	"FOREIGN KEY (DeliveryAddressID) REFERENCES Address(ID) " +
	")"

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx,query)
	return utils.ErrHandler(err)

}
func CreateOrderItemTable (db *sql.DB) error{
	query := `CREATE TABLE IF NOT EXISTS OrderItem (
		ID VARCHAR(255) PRIMARY KEY,
		ProductID VARCHAR(255) NOT NULL,
		OrderID VARCHAR(255) NOT NULL,
		Quantity INT NOT NULL,
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (ProductID) REFERENCES Product(ID),
		FOREIGN KEY (OrderID) REFERENCES ` + "`Order`" + `(ID)
	)`

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx,query)
	return utils.ErrHandler(err)

}