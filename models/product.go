package models

import "time"

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	CategoryID  string `json:"categoryId"`
	SellerID     string `json:"sellerId"`
	Seller       *Seller
	Category    *Category
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
