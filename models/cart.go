package models

import "time"

type Cart struct {
	ID        string     `json:"id"`
	CustomerID    string     `json:"customerId"`
	Items     []CartItem `json:"items"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type CartItem struct {
	ID        string `json:"id"`
	ProductID string `json:"productId"`
	CartID    string `json:"cartId"`
	Product   Product `json:"product"`
	Quantity  int `json:"quantity"`
	CreatedAt         time.Time   `json:"createdAt"`
	UpdatedAt         time.Time   `json:"updatedAt"`
}