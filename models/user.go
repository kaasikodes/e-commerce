package models

import "time"

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"-" validate:"min:6 max:12"`
	Image     string `json:"image"`
	Customer  *Customer `json:"customer"`
	Seller    *Seller `json:"seller"`
	Cart      *Cart `json:"cart"`
	EmailVerified bool `json:"emailVerified"`
	EmailVerifiedAt time.Time `json:"emailVerifiedAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Customer struct {
	ID     string  `json:"id"`
	UserID string  `json:"userId"`
	Orders []Order `json:"orders"`
	CreatedAt         time.Time   `json:"createdAt"`
	UpdatedAt         time.Time   `json:"updatedAt"`
}

type Seller struct {
	ID       string `json:"id"`
	UserID   string `json:"userId"`
	Products []Product `json:"products"`
	CreatedAt         time.Time   `json:"createdAt"`
	UpdatedAt         time.Time   `json:"updatedAt"`
}