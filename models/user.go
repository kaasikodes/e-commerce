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
	EmailVerified bool `json:"emailVerified"`
	EmailVerifiedAt time.Time `json:"emailVerifiedAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Customer struct {
	ID     string  `json:"id"`
	UserID string  `json:"userId"`
	Cart      *Cart `json:"cart"`
	User    *User `json:"user"`
	Orders []Order `json:"orders"`
	CreatedAt         time.Time   `json:"createdAt"`
	UpdatedAt         time.Time   `json:"updatedAt"`
}

type Seller struct {
	ID       string `json:"id"`
	UserID   string `json:"userId"`
	User    *User `json:"user"`
	Products []Product `json:"products"`
	CreatedAt         time.Time   `json:"createdAt"`
	UpdatedAt         time.Time   `json:"updatedAt"`
}