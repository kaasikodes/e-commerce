package migrations

import (
	"context"
	"database/sql"
	"time"

	"github.com/kaasikodes/e-commerce-go/utils"
)

func CreateCategoryTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS Category (
		ID VARCHAR(255) PRIMARY KEY,
		Name VARCHAR(35) NOT NULL,
		Description VARCHAR(100) NOT NULL,
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, query)
	return utils.ErrHandler(err)

}