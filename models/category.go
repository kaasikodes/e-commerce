package models

import "time"

type Category struct {
	ID          string    `json:"id"`
	Name        string    `json:"name" validate:"required,min=3,max=35"`
	Description string    `json:"description" validate:"omitempty,min=3,max=100"`
	Products    []Product `json:"products"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}