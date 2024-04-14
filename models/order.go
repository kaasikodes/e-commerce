package models

import "time"

type Order struct {
	ID                string      `json:"id"`
	CustomerID            string      `json:"customerId"`
	Items             []OrderItem `json:"items"`
	Payment           Payment     `json:"payment"`
	TotalAmount       float64         `json:"totalAmount"`
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
	TotalPrice float64 `json:"totalPrice"`
	CreatedAt         time.Time   `json:"createdAt"`
	UpdatedAt         time.Time   `json:"updatedAt"`
}