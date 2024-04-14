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

type jwtAuthUserContextKey string
type jwtUserIdMapKey string
type PaystackTransactionStatus struct {
	Abandoned string `json:"abandoned"`
	Success string `json:"success"`
}
var (
	PaystackTransactionStatuses = PaystackTransactionStatus{
		Abandoned: "abandoned",
		Success: "success",
	}
)
// default values
const (
	DefaultPageSize 		= 20
	DefaultContextTimeOut 	= time.Second * 5
	FrontendUrl = "http://localhost:3000"
	JWTSecret	= "secret"
	JWTExpirationTime = time.Hour * 24
	PaystackSecretKey = "sk_test_dc0078426d6a4b0cf15b370c15a61de841a23f78"
	PaystackPublicKey = "pk_test_8ad0429e25af1f59ecf24104442f56ee4bbb39fe"
	
	
	
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
	JWTAuthUserContextKey jwtAuthUserContextKey  = "user"
	JWTUserIdMapKey jwtUserIdMapKey = "userID"


)

