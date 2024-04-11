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

func (r *UserRepository) RetrieveSellers(input types.RetrievUsersInput) (types.PaginatedDataOutput, error) {
    db := r.db
    // Prepare query
    query := `
	SELECT 
		s.ID,
		s.UserID,
		s.CreatedAt,
		s.UpdatedAt,
		u.ID AS user_id,
		u.Name AS user_name,
		u.Email AS user_email,
		u.Image AS user_image,
		u.CreatedAt AS user_created_at, 
		u.UpdatedAt AS user_updated_at,
		(SELECT COUNT(*) FROM Seller) AS total_sellers
	FROM 
		Seller s
	JOIN 
		User u ON s.UserID = u.ID
	WHERE 
		s.ID > ? 
	ORDER BY 
		s.ID ASC
	LIMIT ?

    `

    // Create a context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
    defer cancel()

    // Prepare the statement
    output := types.PaginatedDataOutput{}
    stmt, err := db.PrepareContext(ctx, query)
    if err != nil {
        return output, err
    }
    defer stmt.Close()

    // Execute the statement
    sellers := []models.Seller{}
    rows, err := stmt.QueryContext(ctx, utils.Ternary(input.Pagination.NextCursor == "", "", input.Pagination.NextCursor), utils.Ternary(input.Pagination.PageSize == 0, constants.DefaultPageSize, input.Pagination.PageSize))
    if err != nil {
        return output, err
    }
    defer rows.Close()
    total := 0
    for rows.Next() {
        seller := models.Seller{}
		seller.User = &models.User{}
        err = rows.Scan(&seller.ID, &seller.UserID, &seller.CreatedAt, &seller.UpdatedAt, &seller.User.ID, &seller.User.Name, &seller.User.Email, &seller.User.Image, &seller.User.CreatedAt, &seller.User.UpdatedAt, &total)
        if err != nil {
            return output, err
        }
        sellers = append(sellers, seller)
    }
    lastItemId := ""
    // Select last item in the list
    if len(sellers) > 0 {
        lastItem := sellers[len(sellers)-1]
        lastItemId = lastItem.ID
    }

    output = types.PaginatedDataOutput{
        Data:      sellers,
        NextCursor: lastItemId,
        HasMore:    len(sellers) < total,
        Total:      total,
    }
    return output, nil
}

func (r *UserRepository) RetrieveCustomers(input types.RetrievUsersInput) (types.PaginatedDataOutput, error){
	db := r.db
	// prepare query
    query := `
	SELECT 
		s.ID,
		s.UserID,
		s.CreatedAt,
		s.UpdatedAt,
		u.ID AS user_id,
		u.Name AS user_name,
		u.Email AS user_email,
		u.Image AS user_image,
		u.CreatedAt AS user_created_at, 
		u.UpdatedAt AS user_updated_at,
		(SELECT COUNT(*) FROM Seller) AS total_sellers
	FROM 
		Customer s
	JOIN 
		User u ON s.UserID = u.ID
	WHERE 
		s.ID > ? 
	ORDER BY 
		s.ID ASC
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
	customers := []models.Customer{}
	rows, err := stmt.QueryContext(ctx, utils.Ternary(input.Pagination.NextCursor == "", "", input.Pagination.NextCursor), utils.Ternary(input.Pagination.PageSize == 0, constants.DefaultPageSize, input.Pagination.PageSize))
	if err !=nil {
		return output, err
	}
	defer rows.Close()
	total := 0
	for rows.Next() {
		customer := models.Customer{}
		customer.User = &models.User{}
		err = rows.Scan(&customer.ID, &customer.UserID, &customer.CreatedAt, &customer.UpdatedAt, &customer.User.ID, &customer.User.Name, &customer.User.Email, &customer.User.Image, &customer.User.CreatedAt, &customer.User.UpdatedAt, &total)
		if err !=nil {
			return output, err
		}
		customers = append(customers, customer)
	}
	lastItemId := ""
	// select last item in the list
	if len(customers) > 0 {
		lastItem := customers[len(customers)-1]
		lastItemId = lastItem.ID
	}

	output = types.PaginatedDataOutput{
		Data: customers,
		NextCursor: lastItemId,
		HasMore:    len(customers) < total,
		Total:      total,
	}
	return output, nil

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
	query := `SELECT
				u.ID AS user_id,
				u.Name AS user_name,
				u.Email AS user_email,
				u.Image AS user_image,
				u.Password AS user_password,
				u.EmailVerified AS user_email_verified,
				u.CreatedAt AS user_created_at,
				u.UpdatedAt AS user_updated_at,
				c.ID AS customer_id,
				c.UserID AS customer_user_id,
				c.CreatedAt AS customer_created_at,
				c.UpdatedAt AS customer_updated_at,
				s.ID AS seller_id,
				s.UserID AS seller_user_id,
				s.CreatedAt AS seller_created_at,
				s.UpdatedAt AS seller_updated_at
			FROM
				User u
			JOIN
				Customer c ON u.ID = c.UserID
			JOIN
				Seller s ON u.ID = s.UserID
			WHERE
				u.ID = ?
	`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	user := models.User{}
	user.Customer = &models.Customer{}
	user.Seller = &models.Seller{}
	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return user, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	
	 err = stmt.QueryRowContext(ctx,  id).Scan(&user.ID, &user.Name, &user.Email, &user.Image, &user.Password, &user.EmailVerified, &user.CreatedAt, &user.UpdatedAt, &user.Customer.ID, &user.Customer.UserID, &user.Customer.CreatedAt, &user.Customer.UpdatedAt, &user.Seller.ID, &user.Seller.UserID, &user.Seller.CreatedAt, &user.Seller.UpdatedAt)
	 if err !=nil {
		return user, err
	}
	return user, nil

}
func (r *UserRepository) RetrieveUserByEmail(email string) (models.User, error){
	db := r.db
	// prepare query
	query := `SELECT
				u.ID AS user_id,
				u.Name AS user_name,
				u.Email AS user_email,
				u.Image AS user_image,
				u.Password AS user_password,
				u.EmailVerified AS user_email_verified,
				u.CreatedAt AS user_created_at,
				u.UpdatedAt AS user_updated_at,
				c.ID AS customer_id,
				c.UserID AS customer_user_id,
				c.CreatedAt AS customer_created_at,
				c.UpdatedAt AS customer_updated_at,
				s.ID AS seller_id,
				s.UserID AS seller_user_id,
				s.CreatedAt AS seller_created_at,
				s.UpdatedAt AS seller_updated_at
			FROM
				User u
			JOIN
				Customer c ON u.ID = c.UserID
			JOIN
				Seller s ON u.ID = s.UserID
			WHERE
				u.Email = ?
	`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	user := models.User{}
	user.Customer = &models.Customer{}
	user.Seller = &models.Seller{}
	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return user, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	
	 err = stmt.QueryRowContext(ctx,  email).Scan(&user.ID, &user.Name, &user.Email, &user.Image, &user.Password, &user.EmailVerified, &user.CreatedAt, &user.UpdatedAt, &user.Customer.ID, &user.Customer.UserID, &user.Customer.CreatedAt, &user.Customer.UpdatedAt, &user.Seller.ID, &user.Seller.UserID, &user.Seller.CreatedAt, &user.Seller.UpdatedAt)
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