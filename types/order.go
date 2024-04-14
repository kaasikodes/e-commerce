package types

import "github.com/kaasikodes/e-commerce-go/models"

type RetrievOrdersInput struct {
	Pagination Pagination
}


type PaginatedOrdersDataOutput struct {
	Data       []models.Order `json:"data"`
	NextCursor string      `json:"nextCursor"`
	HasMore    bool        `json:"hasMore"`
	Total      int         `json:"total"`
}

type OrderItemInput struct {
	ProductId string `json:"productId" validate:"required"`
	TotalPrice float64 `json:"totalPrice" validate:"required min=0"`
	Quantity int `json:"quantity" validate:"required min=1"`
}

type CreateOrderInput struct {
	TotalAmount float64 `json:"totalAmount" validate:"required min=0"`
	OrderItems  []OrderItemInput
}
type OrderRepository interface {
	CreateOrder(data CreateOrderInput, customerId,addressId string) ( orderId string, error error)
	RetrieveOrder(id string) (models.Order, error)
	RetrieveOrders(input RetrievOrdersInput, customerId string) (PaginatedOrdersDataOutput, error)
	DeleteOrder(id string) ( error)
	
}