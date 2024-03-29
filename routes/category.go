package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaasikodes/e-commerce-go/controllers"
	"github.com/kaasikodes/e-commerce-go/types"
)

type CategoryRoutes struct {
	categoryRepo types.CategoryRepository
}

func NewCategoryRoutes(categoryRepo types.CategoryRepository) *CategoryRoutes {
	return &CategoryRoutes{
		categoryRepo: categoryRepo,
	}
}

func (c *CategoryRoutes) RegisterCategoryRoutes (router *mux.Router){
	controller := controllers.NewCategoryController(c.categoryRepo)
	router.HandleFunc("/categories", controller.GetCategoriesHandler).Methods(http.MethodGet)
	router.HandleFunc("/categories/{id}", controller.GetCategoryHandler).Methods(http.MethodGet)
	router.HandleFunc("/categories/{id}", controller.EditCategoryHandler).Methods(http.MethodPut)
	router.HandleFunc("/categories/{id}", controller.DeleteCategoryHandler).Methods(http.MethodDelete)
	router.HandleFunc("/categories", controller.AddCategoryHandler).Methods(http.MethodPost)
	router.HandleFunc("/categories/bulk/template", controller.GetImportCategoryTemplateHandler).Methods(http.MethodGet)
	router.HandleFunc("/categories/bulk/import", controller.ImportMultipleCategoryHandler).Methods(http.MethodPost)
}