package types

import "github.com/kaasikodes/e-commerce-go/models"

type CreateTokenInput struct {
	Email string `json:"email"`
}


type VerifyTokenInput struct {
	Email     string   `json:"email" validate:"required,email"`
	Token  string   `json:"token" validate:"required"`
}
type TokenRepository interface {
	CreateVerificationToken(input CreateTokenInput) (models.VerificationToken, error)
	DeleteVerificationToken(email string) ( error)
	CreatePasswordResetToken(input CreateTokenInput) (models.PasswordResetToken, error)
	DeletePasswordResetToken(email string) ( error)
	RetrieveVerificationToken(email string) (models.VerificationToken, error)
	RetrievePasswordResetToken(email string) (models.PasswordResetToken, error)
}