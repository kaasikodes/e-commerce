package services

import (
	"context"
	"database/sql"

	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/models"
	"github.com/kaasikodes/e-commerce-go/types"
	"github.com/kaasikodes/e-commerce-go/utils"
)

type TokenRepository struct {
	db *sql.DB
}

func NewTokenRepository(db *sql.DB) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}


func (r *TokenRepository) CreateVerificationToken(input types.CreateTokenInput) (models.VerificationToken, error){
	db := r.db
	// prepare query
	query := `INSERT INTO VerificationToken (ID, Token, Email, ExpiresAt) VALUES (?, ?, ?,?)`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	token := models.VerificationToken{}
	if err !=nil {
		return token, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	id, _ := utils.GenerateRandomID(10)
	tokenVal, _ := utils.GenerateRandomID(14)
	res, err := stmt.ExecContext(ctx, id, tokenVal, input.Email, constants.VerificationTokenExpiresAt)
	if err !=nil {
		return token, err
	}
	_, err = res.RowsAffected()
	if err !=nil {
		return token, err
	}

	token.Email = input.Email
	token.ID = id
	token.Token = tokenVal
	token.ExpiresAt = constants.VerificationTokenExpiresAt

	return token, nil
}
func (r *TokenRepository) DeleteVerificationToken(email string) ( error){
	db := r.db
	// prepare query
	query := "DELETE FROM VerificationToken WHERE Email = ?"
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return  err
	}
	
	
	defer stmt.Close() //close the statement after use

	// execute the statement
	_, err = stmt.ExecContext(ctx, email)
	if err !=nil {
		return  err
	}
	
	
	return  nil
}
func (r *TokenRepository) CreatePasswordResetToken(input types.CreateTokenInput) (models.PasswordResetToken, error){
	db := r.db
	// prepare query
	query := `INSERT INTO PasswordResetToken (ID, Token, Email, ExpiresAt) VALUES (?, ?, ?, ?)`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	token := models.PasswordResetToken{}
	if err !=nil {
		return token, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	id, _ := utils.GenerateRandomID(10)
	tokenVal, _ := utils.GenerateRandomID(14)
	res, err := stmt.ExecContext(ctx, id, tokenVal, input.Email, constants.VerificationTokenExpiresAt)
	if err !=nil {
		return token, err
	}
	_, err = res.RowsAffected()
	if err !=nil {
		return token, err
	}

	token.Email = input.Email
	token.ID = id
	token.Token = tokenVal
	token.ExpiresAt = constants.VerificationTokenExpiresAt

	return token, nil
}
func (r *TokenRepository) DeletePasswordResetToken(email string) ( error){
	db := r.db
	// prepare query
	query := "DELETE FROM PasswordResetToken WHERE Email = ?"
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return  err
	}
	
	
	defer stmt.Close() //close the statement after use

	// execute the statement
	_, err = stmt.ExecContext(ctx, email)
	if err !=nil {
		return  err
	}
	
	
	return  nil
}
func (r *TokenRepository) RetrieveVerificationToken(email string) (models.VerificationToken, error){
	db := r.db
	// prepare query
	query := `SELECT ID,  Email, Token, CreatedAt, ExpiresAt FROM VerificationToken WHERE Email = ?`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	token := models.VerificationToken{}
	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return token, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	
	 err = stmt.QueryRowContext(ctx,  email).Scan(&token.ID, &token.Email, &token.Token, &token.CreatedAt, &token.ExpiresAt)
	 if err !=nil {
		return token, err
	}
	
	return token, nil

}
func (r *TokenRepository) RetrievePasswordResetToken(email string) (models.PasswordResetToken, error){
	db := r.db
	// prepare query
	query := `SELECT ID,  Email, Token, CreatedAt, ExpiresAt FROM PasswordResetToken WHERE Email = ?`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	token := models.PasswordResetToken{}
	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return token, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	
	 err = stmt.QueryRowContext(ctx,  email).Scan(&token.ID, &token.Email, &token.Token, &token.CreatedAt, &token.ExpiresAt)
	 if err !=nil {
		return token, err
	}
	
	return token, nil

}
