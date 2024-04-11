package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaasikodes/e-commerce-go/controllers"
	middleware "github.com/kaasikodes/e-commerce-go/middlware"
	"github.com/kaasikodes/e-commerce-go/types"
)

type UserRoutes struct {
	userRepo types.UserRepository
}

func NewUserRoutes( userRepo types.UserRepository) *UserRoutes {
	return &UserRoutes{
		userRepo: userRepo,
	}
}

func (c *UserRoutes) RegisterUserRoutes (router *mux.Router){
	controller := controllers.NewUserController(c.userRepo)
	middlewareChain := middleware.MiddlewareChain(middleware.RequireAuthMiddleware(c.userRepo))
	
	router.HandleFunc("/users", middlewareChain(controller.GetUsersHandler)).Methods(http.MethodGet)
	router.HandleFunc("/users/customers", middlewareChain(controller.GetCustomersHandler)).Methods(http.MethodGet)
	router.HandleFunc("/users/sellers", middlewareChain(controller.GetSellersHandler)).Methods(http.MethodGet)
	
}