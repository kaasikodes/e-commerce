package services

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/models"
	"github.com/kaasikodes/e-commerce-go/types"
	"github.com/kaasikodes/e-commerce-go/utils"
)
type TRetrievCountriesInput struct {
	Pagination types.Pagination
}
type TRetrievCountriesOutput struct {
	Countries []models.Country
	NextCursor string
	HasMore    bool
	Total      int
}
//  Delete country
func DeleteCountry(db *sql.DB, id string) (models.Country, error) {
	// prepare query
	query := "DELETE FROM Country WHERE ID = ?"
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	utils.ErrHandler(err)
	defer stmt.Close() //close the statement after use
	country, err := RetrieveCountryByID(db, id)
	utils.ErrHandler(err)
	// execute the statement
	_, err = stmt.ExecContext(ctx, id,)
	utils.ErrHandler(err)
	
	
	return country, nil

}
//  Update country by id, with the name & description in data
func UpdateCountry(db *sql.DB, id string, data models.Country) (models.Country, error) {
	// prepare query
	query := "UPDATE country SET Name = ?, Description = ? WHERE ID = ?"
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	utils.ErrHandler(err)
	defer stmt.Close() //close the statement after use
	country, err := RetrieveCountryByID(db, id)
	utils.ErrHandler(err)
	// execute the statement
	res, err := stmt.ExecContext(ctx, id, data.Name,  )
	utils.ErrHandler(err)
	rows, err := res.RowsAffected()
	utils.ErrHandler(err)
	fmt.Println( "These are the rows updated")
	fmt.Println( rows)
	
	fmt.Println( "Updated country")
	fmt.Println( country)
	
	return country, nil

}


// Retrieve countries with pagination
func RetrieveCountries(db *sql.DB, input TRetrievCountriesInput) (TRetrievCountriesOutput, error) {
	// prepare query
	query := `
    SELECT *,
           (SELECT COUNT(*) FROM country) AS total_countries
    FROM country
    WHERE ID > ? 
    ORDER BY ID ASC
    LIMIT ?
	`


	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query )
	utils.ErrHandler(err)
	defer stmt.Close() //close the statement after use
	// execute the statement
	countries := []models.Country{}
	rows, err := stmt.QueryContext(ctx, utils.Ternary(input.Pagination.NextCursor == "", "", input.Pagination.NextCursor), utils.Ternary(input.Pagination.PageSize == 0, constants.DefaultPageSize, input.Pagination.PageSize))
	utils.ErrHandler(err)
	defer rows.Close()
	total := 0
	for rows.Next() {
		country := models.Country{}
		err = rows.Scan(&country.ID, &country.Name,  &total)
		utils.ErrHandler(err)
		countries = append(countries, country)
	}
	lastItemId := ""
	// select last item in the list
	if len(countries) > 0 {
		lastItem := countries[len(countries)-1]
		lastItemId = lastItem.ID
	}

	output := TRetrievCountriesOutput{
		Countries: countries,
		NextCursor: lastItemId,
		HasMore:    len(countries) < total,
		Total:      total,
	}
	return output, nil

}
// Retrieve country by id
func RetrieveCountryByID(db *sql.DB, id string) (models.Country, error) {
	// prepare query
	query := `SELECT * FROM Country WHERE ID = ?`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	utils.ErrHandler(err)
	defer stmt.Close() //close the statement after use
	// execute the statement
	country := models.Country{}
	 err = stmt.QueryRowContext(ctx,  id).Scan(&country.ID, &country.Name)
	utils.ErrHandler(err)
	
	return country, nil

}

//  Add multiple countries
func AddMultipleCountries(db *sql.DB, countries []models.Country) ([]models.Country, error) {
	// prepare query
	query := `INSERT INTO Country (ID, Name, Description) VALUES`
	var inserts []string
	var params []interface{}
	for _, country := range countries {
		inserts = append(inserts, "(?, ?, ?)")
		id, _:= utils.GenerateRandomID(10)
		country.ID = id
		params = append(params, id, country.Name, )

	}
	queryVals := strings.Join(inserts, ",")
	query = query + queryVals
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	utils.ErrHandler(err)
	defer stmt.Close() //close the statement after use
	// execute the statement
	res, err := stmt.ExecContext(ctx, params... )
	utils.ErrHandler(err)
	rows, err := res.RowsAffected()
	fmt.Println( "These are the rows")
	fmt.Println( rows)
	// get the last inserted id
	utils.ErrHandler(err)
	return countries, nil

}
// Add country
func AddCountry(db *sql.DB, c models.Country) (models.Country, error) {
	// prepare query
	query := `INSERT INTO country (ID, Name, Description) VALUES (?, ?, ?)`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	utils.ErrHandler(err)
	defer stmt.Close() //close the statement after use
	// execute the statement
	id, _ := utils.GenerateRandomID(10)
	res, err := stmt.ExecContext(ctx, id, c.Name, )
	utils.ErrHandler(err)
	rows, err := res.RowsAffected()
	fmt.Println( "These are the rows")
	fmt.Println( rows)
	// get the last inserted id
	utils.ErrHandler(err)
	c.ID = id
	return c, nil

}