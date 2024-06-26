package types

import (
	"github.com/kaasikodes/e-commerce-go/models"
)

type CartItemInput struct {
	ProductID string `json:"productId" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required" min:"1"`
}
type SaveCartInput struct {
	Items []CartItemInput `json:"items" validate:"required"`
}
type CartCheckoutInput struct {
	DeliveryAddress AddressInput `json:"deliveryAddress" validate:"required"`
}

type CartRepository interface {
	CreateCart(input SaveCartInput, customerId string) (models.Cart, error)
	DeleteCart(customerId string) error
	RetrieveCart(customerId string) (models.Cart, error)
	CheckoutCart(customerId string, userEmail string) (models.Order, error)
	VerifyPayment(reference string) (VerifyPaystackTransactionResponse, error)
}

