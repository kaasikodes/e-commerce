package types

import "github.com/kaasikodes/e-commerce-go/models"

type RetrievCategoriesInput struct {
	Pagination Pagination
}
type MultipleCategoryInput struct {
	Name        string `json:"name" validate:"required,min=3,max=35"`
	Description string `json:"description" validate:"omitempty,min=3,max=100"`
}

type PaginatedCategoriesDataOutput struct {
	Data       []models.Category `json:"data"`
	NextCursor string      `json:"nextCursor"`
	HasMore    bool        `json:"hasMore"`
	Total      int         `json:"total"`
}

type CategoryRepository interface {
	AddCategory(category models.Category) (models.Category, error)
	UpdateCategory(id string, data models.Category) (models.Category, error)
	RetrieveCategories(pagination RetrievCategoriesInput) (PaginatedCategoriesDataOutput, error)
	RetrieveCategoryByID(id string) (models.Category, error)
	DeleteCategory(id string) (models.Category, error)
	AddMultipleCategories(categories []MultipleCategoryInput) ([]MultipleCategoryInput, error)
}