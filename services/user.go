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

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}


func (r *UserRepository) AddUser(input types.AddUserInput) (models.User, error){
	// ensure that user roles are valid
	validRoles := constants.ValidUserRoles
	for _, role := range input.UserRoles {
		if !strings.Contains(strings.Join(validRoles, ","), role) {
			return models.User{}, constants.ErrInvalidUserRole
		}
	}
	db := r.db
	// prepare query
	query := `INSERT INTO User (ID, Name, Email, Password, Image, Roles) VALUES (?, ?, ?, ?, ?, ?)`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	user := models.User{}
	if err !=nil {
		return user, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	id, _ := utils.GenerateRandomID(10)
	res, err := stmt.ExecContext(ctx, id, input.Name, input.Email, input.Password, input.Image, strings.Join(input.UserRoles, ","))
	if err !=nil {
		return user, err
	}
	_, err = res.RowsAffected()
	if err !=nil {
		return user, err
	}

	user = models.User{ID:id, Name: input.Name, Email: input.Email, Image: input.Image}
	// create the customer and seller records based on the user roles specified
	if(strings.Contains(strings.Join(input.UserRoles, ","), "customer")) {
		customer, err := r.createCustomerProfile(id)
		if err !=nil {
			return user, err
		}
		user.Customer = &customer
	}
	if(strings.Contains(strings.Join(input.UserRoles, ","), "seller")) {
		seller, err := r.createSellerProfile(id)
		if err !=nil {
			return user, err
		}
		user.Seller = &seller
	}

	return user, nil
}
func (r *UserRepository) createCustomerProfile(userId string) (models.Customer, error){
	db := r.db
	// prepare query
	query := `INSERT INTO Customer (ID,UserID) VALUES (?,?)`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	customer := models.Customer{}
	if err !=nil {
		return customer, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	id, _ := utils.GenerateRandomID(10)
	res, err := stmt.ExecContext(ctx, id, userId)
	if err !=nil {
		return customer, err
	}
	_, err = res.RowsAffected()
	if err !=nil {
		return customer, err
	}



	return models.Customer{ID:id, UserID: userId, }, nil
}
func (r *UserRepository) createSellerProfile(userId string) (models.Seller, error){
	db := r.db
	// prepare query
	query := `INSERT INTO Seller (ID,UserID) VALUES (?,?)`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	seller := models.Seller{}
	if err !=nil {
		return seller, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	id, _ := utils.GenerateRandomID(10)
	res, err := stmt.ExecContext(ctx, id, userId)
	if err !=nil {
		return seller, err
	}
	_, err = res.RowsAffected()
	if err !=nil {
		return seller, err
	}



	return models.Seller{ID:id, UserID: userId, }, nil
}
func (r *UserRepository) VerifyUser(email string) (models.User, error){

	db := r.db
	
	// prepare query
	query := "UPDATE User SET EmailVerified = ?, EmailVerifiedAt = ? WHERE Email = ?"
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	user := models.User{}

	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return user, err
	}
	defer stmt.Close() //close the statement after use
	
	res, err := stmt.ExecContext(ctx, true, time.Now(),email )
	if err !=nil {
		return user, err
	}
	_, err = res.RowsAffected()
	if err !=nil {
		return user, err
	}
	user, err = r.RetrieveUserByEmail(email)
	if err !=nil {
		return user, err
	}

	
	return user, nil
}
func (r *UserRepository) UpdateUserPassword(id string, data types.UpdateUserPwdInput) (models.User, error){

	db := r.db
	
	// prepare query
	query := "UPDATE User SET Password = ? WHERE ID = ?"
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	user := models.User{}

	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return user, err
	}
	defer stmt.Close() //close the statement after use
	
	res, err := stmt.ExecContext(ctx, data.Password, id )
	if err !=nil {
		return user, err
	}
	_, err = res.RowsAffected()
	if err !=nil {
		return user, err
	}
	user, err = r.RetrieveUserByID(id)
	if err !=nil {
		return user, err
	}

	
	return user, nil
}
func (r *UserRepository) UpdateUserProfile(id string, data types.UpdateUserProfileInput) (models.User, error){
	db := r.db
	
	// prepare query
	query := "UPDATE User SET Name = ?, Email = ?, Image = ? WHERE ID = ?"
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	user := models.User{}

	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return user, err
	}
	defer stmt.Close() //close the statement after use
	
	res, err := stmt.ExecContext(ctx,  data.Name, data.Email, data.Image, id )
	if err !=nil {
		return user, err
	}
	_, err = res.RowsAffected()
	if err !=nil {
		return user, err
	}
	user, err = r.RetrieveUserByID(id)
	if err !=nil {
		return user, err
	}

	
	return user, nil
}
func (r *UserRepository) RetrieveUsers(input types.RetrievUsersInput) (types.PaginatedDataOutput, error){
	db := r.db
	// prepare query
	query := `
    SELECT ID, Name, Email, Image, CreatedAt, UpdatedAt,
           (SELECT COUNT(*) FROM User) AS total_users
    FROM User
    WHERE ID > ? 
    ORDER BY ID ASC
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
	users := []models.User{}
	rows, err := stmt.QueryContext(ctx, utils.Ternary(input.Pagination.NextCursor == "", "", input.Pagination.NextCursor), utils.Ternary(input.Pagination.PageSize == 0, constants.DefaultPageSize, input.Pagination.PageSize))
	if err !=nil {
		return output, err
	}
	defer rows.Close()
	total := 0
	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Image, &user.CreatedAt, &user.UpdatedAt, &total)
		if err !=nil {
			return output, err
		}
		users = append(users, user)
	}
	lastItemId := ""
	// select last item in the list
	if len(users) > 0 {
		lastItem := users[len(users)-1]
		lastItemId = lastItem.ID
	}

	output = types.PaginatedDataOutput{
		Data: users,
		NextCursor: lastItemId,
		HasMore:    len(users) < total,
		Total:      total,
	}
	return output, nil

}
func (r *UserRepository) RetrieveUserByID(id string) (models.User, error){
	db := r.db
	// prepare query
	query := `SELECT ID, Name, Email, Image, CreatedAt, UpdatedAt FROM User WHERE ID = ?`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	user := models.User{}
	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return user, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	
	 err = stmt.QueryRowContext(ctx,  id).Scan(&user.ID, &user.Name, &user.Email, &user.Image, &user.CreatedAt, &user.UpdatedAt)
	 if err !=nil {
		return user, err
	}
	
	return user, nil

}
func (r *UserRepository) RetrieveUserByEmail(email string) (models.User, error){
	db := r.db
	// prepare query
	query := `SELECT ID, Name, Email, Image, CreatedAt, UpdatedAt FROM User WHERE Email = ?`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	user := models.User{}
	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return user, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	
	 err = stmt.QueryRowContext(ctx,  email).Scan(&user.ID, &user.Name, &user.Email, &user.Image, &user.CreatedAt, &user.UpdatedAt)
	 if err !=nil {
		return user, err
	}
	
	return user, nil

}
func (r *UserRepository) DeleteUser(id string) (models.User, error){
	db := r.db
	// prepare query
	query := "DELETE FROM User WHERE ID = ?"
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	user := models.User{}
	if err !=nil {
		return user, err
	}
	
	defer stmt.Close() //close the statement after use
	user, err = r.RetrieveUserByID(id)
	if err !=nil {
		return user, err
	}
	// execute the statement
	_, err = stmt.ExecContext(ctx, id)
	if err !=nil {
		return user, err
	}
	
	
	return user, nil
}

func (r *UserRepository) AddMultipleUsers(input []types.MultipleUserInput) ([]types.MultipleUserInput, error){
	db := r.db
	
	// prepare query
	query := `INSERT INTO User (ID, Name, Email, Password, Roles) VALUES`
	var inserts []string
	var params []interface{}
	for _, data := range input {
		inserts = append(inserts, "(?, ?, ?, ?, ?)")
		id, _:= utils.GenerateRandomID(10)
		params = append(params, id, data.Name, data.Email, data.Password, strings.Join(data.UserRoles, ","))

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