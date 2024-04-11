package services

import (
	"context"
	"database/sql"
	"strings"

	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/models"
	"github.com/kaasikodes/e-commerce-go/types"
	"github.com/kaasikodes/e-commerce-go/utils"
)

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}
func (c *CartRepository) retrieveCartItems(cartId string) ([]models.CartItem, error) {
	db := c.db
	// prepare query
	query := `
    SELECT *
    FROM CartItem
	WHERE CartID = ?
	`


	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	items := []models.CartItem{}
	stmt, err := db.PrepareContext(ctx, query )
	if err !=nil {
		return items, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	rows, err := stmt.QueryContext(ctx, cartId)
	if err !=nil {
		return items, err
	}
	defer rows.Close()
	for rows.Next() {
		item := models.CartItem{}
		err = rows.Scan(&item.ID, &item.ProductID, &item.CartID, &item.Quantity,  &item.CreatedAt, &item.UpdatedAt)
		if err !=nil {
			return items, err
		}
		items = append(items, item)
	}
	


	return items, nil

}
func (c *CartRepository) createCartItems(input types.SaveCartInput, cartId string) ([]models.CartItem, error) {
	db := c.db
	items := input.Items
	cartItems := []models.CartItem{}
	
	// prepare query
	query := `INSERT INTO CartItem (ID, ProductID, Quantity, CartID) VALUES`
	var inserts []string
	var params []interface{}
	for _, data := range items {
		cartItem := models.CartItem{}
		inserts = append(inserts, "(?, ?, ?, ?)")
		id, _:= utils.GenerateRandomID(10)
		params = append(params, id, data.ProductID, data.Quantity, cartId)
		cartItem.CartID = cartId
		cartItem.ProductID = data.ProductID
		cartItem.Quantity = data.Quantity;
		cartItem.Product = models.Product{}
		cartItems = append(cartItems, cartItem)

	}
	queryVals := strings.Join(inserts, ",")
	query = query + queryVals
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return cartItems, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	res, err := stmt.ExecContext(ctx, params... )
	if err !=nil {
		return cartItems, err
	}
	_, err = res.RowsAffected()
	if err !=nil {
		return cartItems, err
	}
	// ensure input have
	return cartItems, nil

}
func (c *CartRepository) CreateCart(input types.SaveCartInput, customerId string) (models.Cart, error) {
	db := c.db
	// prepare query
	query := `INSERT INTO Cart (ID, CustomerID) VALUES (?, ?)`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	cart := models.Cart{}
	if err != nil {
		return cart, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	id, _ := utils.GenerateRandomID(10)
	_, err = stmt.ExecContext(ctx, id,  customerId)
	if err != nil {
		return  cart, err
	}

	cart.ID = id
	cart.CustomerID = customerId

	cartItems, err := c.createCartItems(input, id)
	if err != nil {
		return  cart, err
	}
	cart.Items = cartItems
	return  cart, nil

}
func (c *CartRepository) RetrieveCart(customerId string) (models.Cart, error){
	db := c.db
	// prepare query
	query := `SELECT * FROM Cart WHERE CustomerID = ?`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	cart := models.Cart{}
	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return cart, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	
	 err = stmt.QueryRowContext(ctx,  customerId).Scan(&cart.ID, &cart.CustomerID, &cart.CreatedAt, &cart.UpdatedAt)
	 if err !=nil {
		return cart, err
	}

	items, err := c.retrieveCartItems(cart.ID)
	if err !=nil {
		return cart, err
	}
	cart.Items = items
	
	return cart, nil

}

// Delete Cart
func (c *CartRepository)DeleteCart(customerId string) (error){
	db := c.db
	// prepare query
	query := "DELETE FROM Cart WHERE CustomerID = ?"
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return err
	}
	
	defer stmt.Close() //close the statement after use
	_, err = c.RetrieveCart(customerId)
	if err !=nil {
		return err
	}
	// execute the statement
	_, err = stmt.ExecContext(ctx, customerId)
	if err !=nil {
		return  err
	}
	
	
	return  nil
}

