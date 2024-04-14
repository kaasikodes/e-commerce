package types

import (
	"time"

	"github.com/kaasikodes/e-commerce-go/models"
)

type RetrievePaymentsInput struct {
	Pagination Pagination
}


type PaginatedPaymentsDataOutput struct {
	Data       []models.Payment `json:"data"`
	NextCursor string      `json:"nextCursor"`
	HasMore    bool        `json:"hasMore"`
	Total      int         `json:"total"`
}

type CreatePaymentInput struct {
	Reference string `json:"reference" validate:"required"`
	Amount float64 `json:"amount" validate:"required min=0"`
	Paid bool `json:"paid"`
	Method string `json:"method" validate:"required"`
	PaidAt time.Time `json:"paidAt"`
	
	
}
type UpdatePaymentInput struct {

	Paid bool `json:"paid"`
	PaidAt time.Time `json:"paidAt"`
	
	
}

type PaymentRepository interface {
	// create payment
// retrieve payment
// retrieve payments
	CreatePayment(data CreatePaymentInput, orderId string) ( error)
	UpdatePayment(data UpdatePaymentInput, reference string) ( error)
	RetrievePayment(id string) (models.Payment, error)
	RetrievePayments(input  RetrievePaymentsInput, customerId string) (PaginatedPaymentsDataOutput, error)
	
}