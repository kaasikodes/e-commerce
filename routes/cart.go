package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaasikodes/e-commerce-go/controllers"
	middleware "github.com/kaasikodes/e-commerce-go/middlware"
	"github.com/kaasikodes/e-commerce-go/types"
)

type CartRoutes struct {
	cartRepo types.CartRepository
	userRepo types.UserRepository
}

func NewCartRoutes(  cartRepo types.CartRepository,  userRepo types.UserRepository) *CartRoutes {
	return &CartRoutes{
		cartRepo: cartRepo,
		userRepo: userRepo,
	}
}

func (c *CartRoutes) RegisterCartRoutes (router *mux.Router){
	controller := controllers.NewCartController( c.cartRepo)
	middlewareChain := middleware.MiddlewareChain(middleware.RequireAuthMiddleware(c.userRepo))
	
	router.HandleFunc("/cart", middlewareChain(controller.SaveCartHandler)).Methods(http.MethodPost)
	router.HandleFunc("/cart", middlewareChain(controller.GetCartHandler)).Methods(http.MethodGet)
	router.HandleFunc("/cart", middlewareChain(controller.DeleteCartHandler)).Methods(http.MethodDelete)

	
	
}