package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaasikodes/e-commerce-go/controllers"
	middleware "github.com/kaasikodes/e-commerce-go/middlware"
	"github.com/kaasikodes/e-commerce-go/types"
)

type ProductRoutes struct {
	userRepo types.UserRepository
	productRepo types.ProductRepository
	categoryRepo types.CategoryRepository
}

func NewProductRoutes( userRepo types.UserRepository, productRepo types.ProductRepository, categoryRepo types.CategoryRepository) *ProductRoutes {
	return &ProductRoutes{
		userRepo: userRepo,
		productRepo: productRepo,
		categoryRepo: categoryRepo,
	}
}

func (c *ProductRoutes) RegisterProductRoutes (router *mux.Router){
	controller := controllers.NewProductController(c.productRepo, c.categoryRepo)
	middlewareChain := middleware.MiddlewareChain(middleware.RequireAuthMiddleware(c.userRepo))
	
	router.HandleFunc("/products", middlewareChain(controller.AddProductHandler)).Methods(http.MethodPost)
	router.HandleFunc("/products", middlewareChain(controller.GetProductsHandler)).Methods(http.MethodGet)
	router.HandleFunc("/products/{id}", middlewareChain(controller.EditProductHandler)).Methods(http.MethodPut)
	router.HandleFunc("/products/{id}", middlewareChain(controller.DeleteProductHandler)).Methods(http.MethodDelete)
	router.HandleFunc("/products/{id}", middlewareChain(controller.GetProductHandler)).Methods(http.MethodGet)
	router.HandleFunc("/products/bulk/template", middlewareChain(controller.GetImportProductTemplateHandler)).Methods(http.MethodGet)
	router.HandleFunc("/products/bulk/import", middlewareChain(controller.ImportMultipleProductHandler)).Methods(http.MethodPost)
	
	
}