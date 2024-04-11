package types

import "github.com/kaasikodes/e-commerce-go/models"

type AddProductInput struct {
	Name string `json:"name" validate:"required,min=3,max=35"`
	Description string `json:"description" validate:"omitempty,min=3,max=100"`
	Price int `json:"price" validate:"required"`
	Quantity int `json:"quantity" validate:"required"`
	CategoryID string `json:"categoryId" validate:"required"`
}
type RetrievProductsInput struct {
	Pagination Pagination
}

type MultipleProductInput struct {
	Name string `json:"name" validate:"required,min=3,max=35"`
	Description string `json:"description" validate:"omitempty,min=3,max=100"`
	Price int `json:"price" validate:"required"`
	Quantity int `json:"quantity" validate:"required"`
	CategoryID string `json:"categoryId" validate:"required"`
	
}
type ProductRepository interface {
	AddProduct(input AddProductInput, sellerId string) (models.Product, error)
	UpdateProduct(id string, input AddProductInput) (models.Product, error)
	AddMultipleProducts(input []MultipleProductInput, sellerId string) ([]MultipleProductInput, error)
	RetrieveProducts(input RetrievProductsInput, sellerId string) (PaginatedDataOutput, error)
	RetrieveProductByID(id string) (models.Product, error)
	DeleteProduct(id string) (models.Product, error)
}

