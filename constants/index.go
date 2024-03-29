package constants

import (
	"errors"
	"time"
)

// config
const (
	DbUser                  = "root"
	DbPassword              = ""
	DbNet                   = "tcp"
	DbName                  = "ecommerce"
	DbHost                  = "127.0.0.1:3306"
	DBMaxConnectionLifeTime = time.Minute * 3
	DBMaxOpenConnections    = 10
	DBMaxIdleConnections    = 10
)
// default values
const (
	DefaultPageSize 		= 20
	DefaultContextTimeOut 	= time.Second * 5
	FrontendUrl = "http://localhost:3000"
	
)
// query keys for urls
const (
	QueryPageSize = "pageSize"
)
// messages
const (
	MsgValidationError = "Validation Error"
	MsgCategoriesRetrieved = "Categories Retrieved Successfully"
	MsgCategoryRetrieved = "Category Retrieved Successfully"
	MsgInternalServerError = "Internal Server Error"
)
// errors
var (
	ErrPageSizeNotValid = errors.New("page size must be an integer")
	ErrCategoryNotFound = errors.New("category not found")
	ErrInvalidUserRole = errors.New("user role should be either customer, or seller")
)
// expirations & general
var (
	VerificationTokenExpiresAt = time.Now().Add(time.Hour * 24)
	PasswordResetTokenExpiresAt = time.Now().Add(time.Hour * 4)
	ValidUserRoles = []string{"customer", "seller"}

)

