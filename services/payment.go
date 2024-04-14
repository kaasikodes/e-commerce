package services

import (
	"context"
	"database/sql"

	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/models"
	"github.com/kaasikodes/e-commerce-go/types"
	"github.com/kaasikodes/e-commerce-go/utils"
)

type PaymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}
// update payment
func (c *PaymentRepository) UpdatePayment(data types.UpdatePaymentInput, reference string) ( error) {
	db := c.db
	// prepare query
	query := "UPDATE Payment SET Paid = ?, PaidAt = ? WHERE ID = ?"
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	res, err := stmt.ExecContext(ctx, data.Paid, data.PaidAt, reference)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil

}
// create payment
func (c *PaymentRepository) CreatePayment(data types.CreatePaymentInput, orderId string) ( error) {
	db := c.db
	// prepare query
	query := `INSERT INTO Payment (ID, OrderID, Amount, Paid, Method) VALUES (?, ?, ?, ?, ?)`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	res, err := stmt.ExecContext(ctx, data.Reference, orderId, data.Amount, data.Paid, data.Method)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil

}
// retrieve payment
func (c *PaymentRepository) RetrievePayment(id string) (models.Payment, error) {

	db := c.db
	payment := models.Payment{}
	// prepare query
	query := `SELECT * FROM Payment WHERE ID = ?`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return payment, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	row := stmt.QueryRowContext(ctx, id)
	err = row.Scan(&payment.ID, &payment.OrderID, &payment.Amount, &payment.Paid, &payment.PaidAt, &payment.Method)
	if err != nil {
		return payment, err
	}

	return payment, nil
	
}
// retrieve payments
func (c *PaymentRepository) RetrievePayments(input types.RetrievePaymentsInput,  customerId string) (types.PaginatedPaymentsDataOutput, error) {
	db := c.db
	// prepare query
	query := `
    SELECT p.*,
           (SELECT COUNT(*) FROM ` + "`Order`" + ` WHERE CustomerID = ?) AS total_payments
    FROM Payment p
    JOIN ` + "`Order`" + ` o ON o.ID = p.OrderID
    WHERE o.CustomerID = ? AND p.ID > ?
    ORDER BY p.ID ASC
    LIMIT ?`


	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	output := types.PaginatedPaymentsDataOutput{}
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return output, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	// execute the statement
	payments := []models.Payment{}
	rows, err := stmt.QueryContext(ctx, customerId, customerId, utils.Ternary(input.Pagination.NextCursor == "", "", input.Pagination.NextCursor), utils.Ternary(input.Pagination.PageSize == 0, constants.DefaultPageSize, input.Pagination.PageSize))
	if err !=nil {
		return output, err
	}
	defer rows.Close()
	total := 0
	for rows.Next() {
		payment := models.Payment{}
		err = rows.Scan(&payment.ID, &payment.OrderID, &payment.Amount, &payment.Paid, &payment.PaidAt, &payment.Method,  &total)
		if err !=nil {
			return output, err
		}
		payments = append(payments, payment)
	}
	lastItemId := ""
	// select last item in the list
	if len(payments) > 0 {
		lastItem := payments[len(payments)-1]
		lastItemId = lastItem.ID
	}

	output.Data = payments
	output.NextCursor = lastItemId
	output.HasMore = len(payments) < total
	output.Total = total
	return output, nil
	
}