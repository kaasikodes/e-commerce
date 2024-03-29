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
type TRetrievStatesInput struct {
	Pagination types.Pagination
}
type TRetrievStatesOutput struct {
	states []models.State
	NextCursor string
	HasMore    bool
	Total      int
}
//  Delete state
func DeleteState(db *sql.DB, id string) (models.State, error) {
	// prepare query
	query := "DELETE FROM state WHERE ID = ?"
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	utils.ErrHandler(err)
	defer stmt.Close() //close the statement after use
	state, err := RetrieveStateByID(db, id)
	utils.ErrHandler(err)
	// execute the statement
	_, err = stmt.ExecContext(ctx, id,)
	utils.ErrHandler(err)
	
	
	return state, nil

}
//  Update state by id, with the name & description in data
func UpdateState(db *sql.DB, id string, data models.State) (models.State, error) {
	// prepare query
	query := "UPDATE state SET Name = ?, Description = ? WHERE ID = ?"
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	utils.ErrHandler(err)
	defer stmt.Close() //close the statement after use
	state, err := RetrieveStateByID(db, id)
	utils.ErrHandler(err)
	// execute the statement
	res, err := stmt.ExecContext(ctx, id, data.Name, data.CountryID )
	utils.ErrHandler(err)
	rows, err := res.RowsAffected()
	utils.ErrHandler(err)
	fmt.Println( "These are the rows updated")
	fmt.Println( rows)
	
	fmt.Println( "Updated state")
	fmt.Println( state)
	
	return state, nil

}


// Retrieve states with pagination
func RetrieveStates(db *sql.DB, input TRetrievStatesInput) (TRetrievStatesOutput, error) {
	// prepare query
	query := `
    SELECT *,
           (SELECT COUNT(*) FROM state) AS total_states
    FROM state
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
	states := []models.State{}
	rows, err := stmt.QueryContext(ctx, utils.Ternary(input.Pagination.NextCursor == "", "", input.Pagination.NextCursor), utils.Ternary(input.Pagination.PageSize == 0, constants.DefaultPageSize, input.Pagination.PageSize))
	utils.ErrHandler(err)
	defer rows.Close()
	total := 0
	for rows.Next() {
		state := models.State{}
		err = rows.Scan(&state.ID, &state.Name, &state.CountryID,  &total)
		utils.ErrHandler(err)
		states = append(states, state)
	}
	lastItemId := ""
	// select last item in the list
	if len(states) > 0 {
		lastItem := states[len(states)-1]
		lastItemId = lastItem.ID
	}

	output := TRetrievStatesOutput{
		states: states,
		NextCursor: lastItemId,
		HasMore:    len(states) < total,
		Total:      total,
	}
	return output, nil

}
// Retrieve state by id
func RetrieveStateByID(db *sql.DB, id string) (models.State, error) {
	// prepare query
	query := `SELECT * FROM state WHERE ID = ?`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	utils.ErrHandler(err)
	defer stmt.Close() //close the statement after use
	// execute the statement
	state := models.State{}
	 err = stmt.QueryRowContext(ctx,  id).Scan(&state.ID, &state.Name, &state.CountryID)
	utils.ErrHandler(err)
	
	return state, nil

}

//  Add multiple states
func AddMultipleStates(db *sql.DB, states []models.State) ([]models.State, error) {
	// prepare query
	query := `INSERT INTO state (ID, Name, Description) VALUES`
	var inserts []string
	var params []interface{}
	for _, state := range states {
		inserts = append(inserts, "(?, ?, ?)")
		id, _:= utils.GenerateRandomID(10)
		state.ID = id
		params = append(params, id, state.Name, state.CountryID )

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
	return states, nil

}
// Add state
func AddState(db *sql.DB, c models.State) (models.State, error) {
	// prepare query
	query := `INSERT INTO state (ID, Name, Description) VALUES (?, ?, ?)`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	utils.ErrHandler(err)
	defer stmt.Close() //close the statement after use
	// execute the statement
	id, _ := utils.GenerateRandomID(10)
	res, err := stmt.ExecContext(ctx, id, c.Name, c.CountryID)
	utils.ErrHandler(err)
	rows, err := res.RowsAffected()
	fmt.Println( "These are the rows")
	fmt.Println( rows)
	// get the last inserted id
	utils.ErrHandler(err)
	c.ID = id
	return c, nil

}