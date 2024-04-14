package services

import (
	"context"
	"database/sql"

	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/models"
	"github.com/kaasikodes/e-commerce-go/types"
	"github.com/kaasikodes/e-commerce-go/utils"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

// create order
func (c *OrderRepository) CreateOrder(data types.CreateOrderInput, customerId, addressId string) ( orderId string, error error) {
	db := c.db
	orderId = ""
	// prepare query
	query := `INSERT INTO ` + "`Order`" + ` (ID, CustomerID, TotalAmount, DeliveryAddressID) VALUES (?, ?, ?, ?)`

	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return orderId, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	orderId, _ = utils.GenerateRandomID(10)
	res, err := stmt.ExecContext(ctx, orderId, customerId, data.TotalAmount, addressId)
	if err != nil {
		return orderId, err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return orderId, err
	}

	// create order items
	if err = c.createOrderItems(orderId, data.OrderItems); err != nil {
		return orderId, err
	}

	return orderId,nil
}
// retrieve order
func (c *OrderRepository) RetrieveOrder(id string) (models.Order, error) {

	db := c.db
	order := models.Order{}
	// prepare query
	query := `SELECT * FROM ` + "`Order`" + ` o JOIN Payment p ON o.ID = p.OrderID WHERE o.ID = ?`

	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return order, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	row := stmt.QueryRowContext(ctx, id)
	order.Payment = models.Payment{}
	err = row.Scan(&order.ID, &order.CustomerID, &order.TotalAmount, &order.DeliveryAddressID, &order.CreatedAt, &order.UpdatedAt, &order.Payment.ID, &order.Payment.OrderID, &order.Payment.Amount, &order.Payment.Paid, &order.Payment.PaidAt, &order.Payment.Method)
	if err != nil {
		return order, err
	}

	// retrieve order items
	orderItems, err := c.retrieveOrderItems(order.ID)
	if err != nil {
		return order, err
	}
	order.Items = orderItems
	return order, nil
	
}
// retrieve orders
func (c *OrderRepository) RetrieveOrders(input types.RetrievOrdersInput, customerId string) (types.PaginatedOrdersDataOutput, error)  {
	db := c.db
	// prepare query
	query := `
    SELECT o.*,
           p.ID AS payment_id,
           p.OrderID AS payment_order_id,
           p.Amount AS payment_amount,
           p.Paid AS payment_paid,
           p.PaidAt AS payment_paid_at,
           p.Method AS payment_method,
           (SELECT COUNT(*) FROM ` + "`Order`" + ` WHERE CustomerID = ?) AS total_orders
    FROM ` + "`Order`" + ` o
    JOIN Payment p ON p.OrderID = o.ID
    WHERE o.CustomerID = ? AND o.ID > ?
    ORDER BY o.ID ASC
    LIMIT ?`


	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	output := types.PaginatedOrdersDataOutput{}
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return output, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	// execute the statement
	orders := []models.Order{}
	rows, err := stmt.QueryContext(ctx, customerId, customerId, utils.Ternary(input.Pagination.NextCursor == "", "", input.Pagination.NextCursor), utils.Ternary(input.Pagination.PageSize == 0, constants.DefaultPageSize, input.Pagination.PageSize))
	if err !=nil {
		return output, err
	}
	defer rows.Close()
	total := 0
	for rows.Next() {
		order := models.Order{}
		order.Payment = models.Payment{}
		err = rows.Scan(&order.ID, &order.CustomerID, &order.TotalAmount, &order.DeliveryAddressID,&order.CreatedAt, &order.UpdatedAt, &order.Payment.ID, &order.Payment.OrderID, &order.Payment.Amount, &order.Payment.Paid, &order.Payment.PaidAt, &order.Payment.Method, &total)
		if err !=nil {
			return output, err
		}
		orders = append(orders, order)
	}
	lastItemId := ""
	// select last item in the list
	if len(orders) > 0 {
		lastItem := orders[len(orders)-1]
		lastItemId = lastItem.ID
	}

	output.Data = orders
	output.NextCursor = lastItemId
	output.HasMore = len(orders) < total
	output.Total = total
	return output, nil
	
}
// delete order
func (c *OrderRepository) DeleteOrder(id string) ( error){
	db := c.db
	// prepare query
	query := "DELETE FROM `Order` WHERE ID = ?"

	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return  err
	}
	
	defer stmt.Close() //close the statement after use
	_, err = c.RetrieveOrder(id)
	if err !=nil {
		return  err
	}
	// execute the statement
	_, err = stmt.ExecContext(ctx, id)
	if err !=nil {
		return  err
	}
	
	
	return  nil
}

// private
// create order items
func (c *OrderRepository) createOrderItems(orderId string, items []types.OrderItemInput) error {
	db := c.db
	// prepare query
	query := `INSERT INTO OrderItem (ID, OrderID, ProductID, Quantity, TotalPrice) VALUES (?, ?, ?, ?, ?)`
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
	for _, item := range items {
		id, _ := utils.GenerateRandomID(10)
		_, err := stmt.ExecContext(ctx, id, orderId, item.ProductId, item.Quantity, item.TotalPrice)
		if err != nil {
			return err
		} 
	} 
	return nil 
}

// retieve order items
func (c *OrderRepository) retrieveOrderItems(orderId string) ([]models.OrderItem, error) {
	db := c.db
	// prepare query
	query := `SELECT * FROM OrderItem WHERE OrderID = ?`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	rows, err := stmt.QueryContext(ctx, orderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close() //close the rows after use
	orderItems := []models.OrderItem{}
	for rows.Next() {
		orderItem := models.OrderItem{}
		err = rows.Scan(&orderItem.ID, &orderItem.ProductID, &orderItem.OrderID, &orderItem.Quantity, &orderItem.TotalPrice, &orderItem.CreatedAt, &orderItem.UpdatedAt)
		if err != nil {
			return nil, err
		}
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}