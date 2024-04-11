package types

import "github.com/kaasikodes/e-commerce-go/models"

type ForgotPwdInput struct {
	Email     string   `json:"email" validate:"required,email"`
}
type ResetPwdInput struct {
	Email     string   `json:"email" validate:"required,email"`
	Token  string   `json:"token" validate:"required"`
}
type LoginUserInput struct {
	Email     string   `json:"email" validate:"required,email"`
	Password  string   `json:"password" validate:"required"`
}

type AddUserInput struct {
	Name      string   `json:"name" validate:"required,min=3,max=35"`
	Email     string   `json:"email" validate:"required,email"`
	Password  string   `validate:"min=6,max=12"`
	Image     string   `json:"image" validate:"omitempty,url"`
	UserRoles []string `json:"roles" validate:"required"` //customer or seller or both
}
type UpdateUserPwdInput struct {
	Password string `validate:"min=6,max=12"`
}

type UpdateUserProfileInput struct{
	Name string `json:"name" validate:"required,min=3,max=35"`
	Email string `json:"email" validate:"required,email"`
	Image string `json:"image" validate:"nonempty,url"`
}
type RetrievUsersInput struct {
	Pagination Pagination
}

type MultipleUserInput struct {
	Name string `json:"name" validate:"required,min=3,max=35"`
	Email string `json:"email" validate:"required,email"`
	Password string 
	UserRoles []string `json:"userRoles" validate:"required,oneof=[customer,seller] [seller] [customer]"` //customer or seller or both

}
// repository interfaces
type UserRepository interface {
	AddUser(input AddUserInput) (models.User, error)
	UpdateUserPassword(id string, data UpdateUserPwdInput) (models.User, error)
	VerifyUser(email string) (models.User, error)
	UpdateUserProfile(id string, data UpdateUserProfileInput) (models.User, error)
	RetrieveUsers(input RetrievUsersInput) (PaginatedDataOutput, error)
	RetrieveCustomers(input RetrievUsersInput) (PaginatedDataOutput, error)
	RetrieveSellers(input RetrievUsersInput) (PaginatedDataOutput, error)
	RetrieveUserByEmail(email string) (models.User, error)
	RetrieveUserByID(id string) (models.User, error)
	DeleteUser(id string) (models.User, error)
	AddMultipleUsers(input []MultipleUserInput) ([]MultipleUserInput, error)
}