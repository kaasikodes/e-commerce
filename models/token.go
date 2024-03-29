package models

import "time"

type VerificationToken struct {
	ID          string    `json:"id"`
	Email        string    `json:"email"`
	Token string    `json:"token"`
	CreatedAt   time.Time `json:"createdAt"`
	ExpiresAt   time.Time `json:"expiresAt"`
}
type PasswordResetToken struct {
	ID          string    `json:"id"`
	Email        string    `json:"email"`
	Token string    `json:"token"`
	CreatedAt   time.Time `json:"createdAt"`
	ExpiresAt   time.Time `json:"expiresAt"`
}