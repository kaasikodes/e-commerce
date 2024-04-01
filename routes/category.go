package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaasikodes/e-commerce-go/controllers"
	middleware "github.com/kaasikodes/e-commerce-go/middlware"
	"github.com/kaasikodes/e-commerce-go/types"
)

type CategoryRoutes struct {
	categoryRepo types.CategoryRepository
	userRepo types.UserRepository
}

func NewCategoryRoutes(categoryRepo types.CategoryRepository, userRepo types.UserRepository) *CategoryRoutes {
	return &CategoryRoutes{
		categoryRepo: categoryRepo,
		userRepo: userRepo,
	}
}

func (c *CategoryRoutes) RegisterCategoryRoutes (router *mux.Router){
	controller := controllers.NewCategoryController(c.categoryRepo)
	middlewareChain := middleware.MiddlewareChain(middleware.RequireAuthMiddleware(c.userRepo))
	
	router.HandleFunc("/categories", middlewareChain(controller.GetCategoriesHandler)).Methods(http.MethodGet)
	router.HandleFunc("/categories/{id}", middlewareChain(controller.GetCategoryHandler)).Methods(http.MethodGet)
	router.HandleFunc("/categories/{id}", middlewareChain(controller.EditCategoryHandler)).Methods(http.MethodPut)
	router.HandleFunc("/categories/{id}", middlewareChain(controller.DeleteCategoryHandler)).Methods(http.MethodDelete)
	router.HandleFunc("/categories", middlewareChain(controller.AddCategoryHandler)).Methods(http.MethodPost)
	router.HandleFunc("/categories/bulk/template", middlewareChain(controller.GetImportCategoryTemplateHandler)).Methods(http.MethodGet)
	router.HandleFunc("/categories/bulk/import", middlewareChain(controller.ImportMultipleCategoryHandler)).Methods(http.MethodPost)
}