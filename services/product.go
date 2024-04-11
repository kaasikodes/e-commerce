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

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

// update product
func (c *ProductRepository) UpdateProduct(id string, input types.AddProductInput) (models.Product, error){
	db := c.db
	// prepare query
	query := "UPDATE Product SET Name = ?, Description = ?, Price = ?, Quantity = ?, CategoryID = ? WHERE ID = ?"
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	product := models.Product{}

	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return product, err
	}
	defer stmt.Close() //close the statement after use
	
	res, err := stmt.ExecContext(ctx, input.Name, input.Description, input.Price, input.Quantity, input.CategoryID, id )
	if err !=nil {
		return product, err
	}
	_, err = res.RowsAffected()
	if err !=nil {
		return product, err
	}
	product, err = c.RetrieveProductByID( id)
	if err !=nil {
		return product, err
	}

	
	return product, nil

}

// add multiple products
func (c *ProductRepository) AddMultipleProducts(input []types.MultipleProductInput, sellerId string) ([]types.MultipleProductInput, error){
	db := c.db
	
	// prepare query
	query := `INSERT INTO Product (ID, Name, Description, Price, Quantity, CategoryID, OwnerID) VALUES`
	var inserts []string
	var params []interface{}
	for _, data := range input {
		inserts = append(inserts, "(?, ?, ?, ?, ?, ?, ?)")
		id, _:= utils.GenerateRandomID(10)
		params = append(params, id, data.Name, data.Description, data.Price, data.Quantity, data.CategoryID, sellerId)

	}
	queryVals := strings.Join(inserts, ",")
	query = query + queryVals
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return input, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	res, err := stmt.ExecContext(ctx, params... )
	if err !=nil {
		return input, err
	}
	_, err = res.RowsAffected()
	if err !=nil {
		return input, err
	}
	// ensure input have
	return input, nil

}

// retrieve products
func (c *ProductRepository) RetrieveProducts(input types.RetrievProductsInput, sellerId string) (types.PaginatedDataOutput, error){
	db := c.db
	// prepare query
	query := `
    SELECT p.*,
			c.ID AS category_id,
			c.Name AS category_name,
			c.Description AS category_description,
			c.CreatedAt AS category_created_at,
			c.UpdatedAt AS category_updated_at,
		
           (SELECT COUNT(*) FROM Product WHERE OwnerID = ?) AS total_products
    FROM Product p
	JOIN Category c ON p.CategoryID = c.ID
	WHERE p.OwnerID = ? AND  p.ID > ? 
    ORDER BY p.ID ASC
    LIMIT ?
	`


	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	output := types.PaginatedDataOutput{}
	stmt, err := db.PrepareContext(ctx, query )
	if err !=nil {
		return output, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	products := []models.Product{}
	rows, err := stmt.QueryContext(ctx, sellerId, sellerId, utils.Ternary(input.Pagination.NextCursor == "", "", input.Pagination.NextCursor), utils.Ternary(input.Pagination.PageSize == 0, constants.DefaultPageSize, input.Pagination.PageSize))
	if err !=nil {
		return output, err
	}
	defer rows.Close()
	total := 0
	for rows.Next() {
		product := models.Product{}
		product.Category = &models.Category{}
		product.Seller = &models.Seller{}
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.CategoryID, &product.SellerID, &product.CreatedAt, &product.UpdatedAt, &product.Category.ID, &product.Category.Name, &product.Category.Description, &product.Category.CreatedAt, &product.Category.UpdatedAt, &total)
		if err !=nil {
			return output, err
		}
		products = append(products, product)
	}
	lastItemId := ""
	// select last item in the list
	if len(products) > 0 {
		lastItem := products[len(products)-1]
		lastItemId = lastItem.ID
	}

	output = types.PaginatedDataOutput{
		Data: products,
		NextCursor: lastItemId,
		HasMore:    len(products) < total,
		Total:      total,
	}
	return output, nil

}
// retrieve product by id
func (c *ProductRepository) RetrieveProductByID(id string) (models.Product, error){
	db := c.db
	// prepare query
	query := `SELECT * FROM Product WHERE ID = ?`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	product := models.Product{}
	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return product, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	
	 err = stmt.QueryRowContext(ctx,  id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.CategoryID, &product.SellerID, &product.CreatedAt, &product.UpdatedAt)
	 if err !=nil {
		return product, err
	}
	
	return product, nil

}
// Delete Product
func (c *ProductRepository)DeleteProduct(id string) (models.Product, error){
	db := c.db
	// prepare query
	query := "DELETE FROM Product WHERE ID = ?"
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	product := models.Product{}
	if err !=nil {
		return product, err
	}
	
	defer stmt.Close() //close the statement after use
	product, err = c.RetrieveProductByID(id)
	if err !=nil {
		return product, err
	}
	// execute the statement
	_, err = stmt.ExecContext(ctx, id)
	if err !=nil {
		return product, err
	}
	
	
	return product, nil
}
// Add product
func (c *ProductRepository) AddProduct(inp types.AddProductInput, sellerId string) (models.Product, error) {
	db := c.db
	// prepare query
	query := `INSERT INTO Product (ID, Name, Description, Price, Quantity, CategoryID, OwnerID) VALUES (?, ?, ?, ?, ?, ?, ?)`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	product := models.Product{}
	if err != nil {
		return product, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	id, _ := utils.GenerateRandomID(10)
	res, err := stmt.ExecContext(ctx, id, inp.Name, inp.Description, inp.Price, inp.Quantity, inp.CategoryID, sellerId)
	if err != nil {
		return product, err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return product, err
	}
	product.ID = id
	product.Name = inp.Name
	product.Description = inp.Description
	product.Price = inp.Price
	product.Quantity = inp.Quantity
	product.CategoryID = inp.CategoryID
	product.SellerID = sellerId
	return product, nil

}






