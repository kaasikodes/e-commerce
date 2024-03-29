package models

import "time"

type Order struct {
	ID                string      `json:"id"`
	UserID            string      `json:"userId"`
	Items             []OrderItem `json:"items"`
	Payment           Payment     `json:"payment"`
	TotalAmount       int         `json:"totalAmount"`
	DeliveryAddressID string      `json:"deliveryAddressId"`
	DeliveryAddress   Address     `json:"deliveryAddress"`
	CreatedAt         time.Time   `json:"createdAt"`
	UpdatedAt         time.Time   `json:"updatedAt"`
}

type OrderItem struct {
	ID        string `json:"id"`
	ProductID string `json:"productId"`
	OrderID   string `json:"orderId"`
	Product   Product
	Quantity  int `json:"quantity"`
	CreatedAt         time.Time   `json:"createdAt"`
	UpdatedAt         time.Time   `json:"updatedAt"`
}