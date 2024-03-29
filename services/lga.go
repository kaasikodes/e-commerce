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
type TRetrieveLgaInput struct {
	Pagination types.Pagination
}
type TRetrievLgasOutput struct {
	lgas []models.Lga
	NextCursor string
	HasMore    bool
	Total      int
}
//  Delete lga
func DeleteLga(db *sql.DB, id string) (models.Lga, error) {
	// prepare query
	query := "DELETE FROM lga WHERE ID = ?"
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// prepare the lgament
	stmt, err := db.PrepareContext(ctx, query)
	utils.ErrHandler(err)
	defer stmt.Close() //close the lgament after use
	lga, err := RetrieveLgaByID(db, id)
	utils.ErrHandler(err)
	// execute the lgament
	_, err = stmt.ExecContext(ctx, id,)
	utils.ErrHandler(err)
	
	
	return lga, nil

}
//  Update lga by id, with the name & description in data
func UpdateLga(db *sql.DB, id string, data models.Lga) (models.Lga, error) {
	// prepare query
	query := "UPDATE lga SET Name = ?, Description = ? WHERE ID = ?"
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// prepare the lgament
	stmt, err := db.PrepareContext(ctx, query)
	utils.ErrHandler(err)
	defer stmt.Close() //close the lgament after use
	lga, err := RetrieveLgaByID(db, id)
	utils.ErrHandler(err)
	// execute the lgament
	res, err := stmt.ExecContext(ctx, id, data.Name, data.StateId )
	utils.ErrHandler(err)
	rows, err := res.RowsAffected()
	utils.ErrHandler(err)
	fmt.Println( "These are the rows updated")
	fmt.Println( rows)
	
	fmt.Println( "Updated lga")
	fmt.Println( lga)
	
	return lga, nil

}


// Retrieve lgas with pagination
func RetrieveLgas(db *sql.DB, input TRetrieveLgaInput) (TRetrievLgasOutput, error) {
	// prepare query
	query := `
    SELECT *,
           (SELECT COUNT(*) FROM lga) AS total_lgas
    FROM lga
    WHERE ID > ? 
    ORDER BY ID ASC
    LIMIT ?
	`


	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// prepare the lgament
	stmt, err := db.PrepareContext(ctx, query )
	utils.ErrHandler(err)
	defer stmt.Close() //close the lgament after use
	// execute the lgament
	lgas := []models.Lga{}
	rows, err := stmt.QueryContext(ctx, utils.Ternary(input.Pagination.NextCursor == "", "", input.Pagination.NextCursor), utils.Ternary(input.Pagination.PageSize == 0, constants.DefaultPageSize, input.Pagination.PageSize))
	utils.ErrHandler(err)
	defer rows.Close()
	total := 0
	for rows.Next() {
		lga := models.Lga{}
		err = rows.Scan(&lga.ID, &lga.Name, &lga.StateId,  &total)
		utils.ErrHandler(err)
		lgas = append(lgas, lga)
	}
	lastItemId := ""
	// select last item in the list
	if len(lgas) > 0 {
		lastItem := lgas[len(lgas)-1]
		lastItemId = lastItem.ID
	}

	output := TRetrievLgasOutput{
		lgas: lgas,
		NextCursor: lastItemId,
		HasMore:    len(lgas) < total,
		Total:      total,
	}
	return output, nil

}
// Retrieve lga by id
func RetrieveLgaByID(db *sql.DB, id string) (models.Lga, error) {
	// prepare query
	query := `SELECT * FROM lga WHERE ID = ?`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// prepare the lgament
	stmt, err := db.PrepareContext(ctx, query)
	utils.ErrHandler(err)
	defer stmt.Close() //close the lgament after use
	// execute the lgament
	lga := models.Lga{}
	 err = stmt.QueryRowContext(ctx,  id).Scan(&lga.ID, &lga.Name, &lga.StateId)
	utils.ErrHandler(err)
	
	return lga, nil

}

//  Add multiple lgas
func AddMultipleLgas(db *sql.DB, lgas []models.Lga) ([]models.Lga, error) {
	// prepare query
	query := `INSERT INTO lga (ID, Name, Description) VALUES`
	var inserts []string
	var params []interface{}
	for _, lga := range lgas {
		inserts = append(inserts, "(?, ?, ?)")
		id, _:= utils.GenerateRandomID(10)
		lga.ID = id
		params = append(params, id, lga.Name, lga.StateId )

	}
	queryVals := strings.Join(inserts, ",")
	query = query + queryVals
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// prepare the lgament
	stmt, err := db.PrepareContext(ctx, query)
	utils.ErrHandler(err)
	defer stmt.Close() //close the lgament after use
	// execute the lgament
	res, err := stmt.ExecContext(ctx, params... )
	utils.ErrHandler(err)
	rows, err := res.RowsAffected()
	fmt.Println( "These are the rows")
	fmt.Println( rows)
	// get the last inserted id
	utils.ErrHandler(err)
	return lgas, nil

}
// Add lga
func AddLga(db *sql.DB, c models.Lga) (models.Lga, error) {
	// prepare query
	query := `INSERT INTO lga (ID, Name, Description) VALUES (?, ?, ?)`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// prepare the lgament
	stmt, err := db.PrepareContext(ctx, query)
	utils.ErrHandler(err)
	defer stmt.Close() //close the lgament after use
	// execute the lgament
	id, _ := utils.GenerateRandomID(10)
	res, err := stmt.ExecContext(ctx, id, c.Name, c.StateId)
	utils.ErrHandler(err)
	rows, err := res.RowsAffected()
	fmt.Println( "These are the rows")
	fmt.Println( rows)
	// get the last inserted id
	utils.ErrHandler(err)
	c.ID = id
	return c, nil

}