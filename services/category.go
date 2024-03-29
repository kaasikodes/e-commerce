package services

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/models"
	"github.com/kaasikodes/e-commerce-go/types"
	"github.com/kaasikodes/e-commerce-go/utils"
)


type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}



//  Delete category
func (c *CategoryRepository) DeleteCategory( id string) (models.Category, error) {
	db := c.db
	// prepare query
	query := "DELETE FROM Category WHERE ID = ?"
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	category := models.Category{}
	if err !=nil {
		return category, err
	}
	
	defer stmt.Close() //close the statement after use
	category, err = c.RetrieveCategoryByID(id)
	if err !=nil {
		return category, err
	}
	// execute the statement
	_, err = stmt.ExecContext(ctx, id)
	if err !=nil {
		return category, err
	}
	
	
	return category, nil
}
//  Update category by id, with the name & description in data
func (c *CategoryRepository) UpdateCategory( id string, data models.Category) (models.Category, error) {
	db := c.db
	// prepare query
	query := "UPDATE Category SET Name = ?, Description = ? WHERE ID = ?"
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	category := models.Category{}

	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return category, err
	}
	defer stmt.Close() //close the statement after use
	
	res, err := stmt.ExecContext(ctx,  data.Name, data.Description, id )
	if err !=nil {
		return category, err
	}
	_, err = res.RowsAffected()
	if err !=nil {
		return category, err
	}
	category, err = c.RetrieveCategoryByID( id)
	if err !=nil {
		return category, err
	}

	
	return category, nil

}

// Retrieve categories with pagination
func (c *CategoryRepository) RetrieveCategories(input types.RetrievCategoriesInput) (types.PaginatedDataOutput, error) {
	db := c.db
	// prepare query
	query := `
    SELECT *,
           (SELECT COUNT(*) FROM Category) AS total_categories
    FROM Category
    WHERE ID > ? 
    ORDER BY ID ASC
    LIMIT ?
	`


	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// prepare the statement
	output := types.PaginatedDataOutput{}
	stmt, err := db.PrepareContext(ctx, query )
	if err !=nil {
		return output, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	categories := []models.Category{}
	rows, err := stmt.QueryContext(ctx, utils.Ternary(input.Pagination.NextCursor == "", "", input.Pagination.NextCursor), utils.Ternary(input.Pagination.PageSize == 0, constants.DefaultPageSize, input.Pagination.PageSize))
	if err !=nil {
		return output, err
	}
	defer rows.Close()
	total := 0
	for rows.Next() {
		category := models.Category{}
		err = rows.Scan(&category.ID, &category.Name, &category.Description, &category.CreatedAt, &category.UpdatedAt, &total)
		if err !=nil {
			return output, err
		}
		categories = append(categories, category)
	}
	lastItemId := ""
	// select last item in the list
	if len(categories) > 0 {
		lastItem := categories[len(categories)-1]
		lastItemId = lastItem.ID
	}

	output = types.PaginatedDataOutput{
		Data: categories,
		NextCursor: lastItemId,
		HasMore:    len(categories) < total,
		Total:      total,
	}
	return output, nil

}
// Retrieve category by id
func (c *CategoryRepository) RetrieveCategoryByID( id string) (models.Category, error) {
	db := c.db
	// prepare query
	query := `SELECT * FROM Category WHERE ID = ?`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// prepare the statement
	category := models.Category{}
	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return category, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	
	 err = stmt.QueryRowContext(ctx,  id).Scan(&category.ID, &category.Name, &category.Description, &category.CreatedAt, &category.UpdatedAt)
	 if err !=nil {
		return category, err
	}
	
	return category, nil

}

//  Add multiple categories
func (c *CategoryRepository) AddMultipleCategories( categories []types.MultipleCategoryInput) ([]types.MultipleCategoryInput, error) {
	db := c.db
	
	// prepare query
	query := `INSERT INTO Category (ID, Name, Description) VALUES`
	var inserts []string
	var params []interface{}
	for _, category := range categories {
		inserts = append(inserts, "(?, ?, ?)")
		id, _:= utils.GenerateRandomID(10)
		params = append(params, id, category.Name, category.Description)

	}
	queryVals := strings.Join(inserts, ",")
	query = query + queryVals
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return categories, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	res, err := stmt.ExecContext(ctx, params... )
	if err !=nil {
		return categories, err
	}
	_, err = res.RowsAffected()
	if err !=nil {
		return categories, err
	}
	// ensure categories have
	return categories, nil

}
// Add category
func (c *CategoryRepository) AddCategory( cat models.Category) (models.Category, error) {
	db := c.db
	// prepare query
	query := `INSERT INTO Category (ID, Name, Description) VALUES (?, ?, ?)`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	category := models.Category{}
	if err !=nil {
		return category, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	id, _ := utils.GenerateRandomID(10)
	res, err := stmt.ExecContext(ctx, id, cat.Name, cat.Description, )
	if err !=nil {
		return category, err
	}
	_, err = res.RowsAffected()
	if err !=nil {
		return category, err
	}
	cat.ID = id
	return cat, nil

}