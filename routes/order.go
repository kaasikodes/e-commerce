package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaasikodes/e-commerce-go/controllers"
	middleware "github.com/kaasikodes/e-commerce-go/middlware"
	"github.com/kaasikodes/e-commerce-go/types"
)

type OrderRoutes struct {
	orderRepo types.OrderRepository
	userRepo types.UserRepository
}

func NewOrderRoutes(  orderRepo types.OrderRepository,  userRepo types.UserRepository) *OrderRoutes {
	return &OrderRoutes{
		orderRepo: orderRepo,
		userRepo: userRepo,
	}
}

func (c *OrderRoutes) RegisterOrderRoutes (router *mux.Router){
	controller := controllers.NewOrderController( c.orderRepo,  c.userRepo)
	middlewareChain := middleware.MiddlewareChain(middleware.RequireAuthMiddleware(c.userRepo))
	
	router.HandleFunc("/orders/{id}", middlewareChain(controller.DeleteOrderHandler)).Methods(http.MethodDelete)
	router.HandleFunc("/orders/{id}", middlewareChain(controller.GetOrderHandler)).Methods(http.MethodGet)
	router.HandleFunc("/orders", middlewareChain(controller.GetOrdersHandler)).Methods(http.MethodGet)


	
	
}