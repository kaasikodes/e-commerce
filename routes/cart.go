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
	orderRepo types.OrderRepository
	paymentRepo types.PaymentRepository
	addressRepo types.AddressRepository
}

func NewCartRoutes(  cartRepo types.CartRepository,  userRepo types.UserRepository, orderRepo types.OrderRepository, paymentRepo types.PaymentRepository, addressRepo types.AddressRepository) *CartRoutes {
	return &CartRoutes{
		cartRepo: cartRepo,
		userRepo: userRepo,
		orderRepo: orderRepo,
		paymentRepo: paymentRepo,
		addressRepo: addressRepo,
	}
}

func (c *CartRoutes) RegisterCartRoutes (router *mux.Router){
	controller := controllers.NewCartController(c.cartRepo, c.orderRepo, c.paymentRepo, c.addressRepo)
	middlewareChain := middleware.MiddlewareChain(middleware.RequireAuthMiddleware(c.userRepo))
	
	router.HandleFunc("/cart", middlewareChain(controller.SaveCartHandler)).Methods(http.MethodPost)
	router.HandleFunc("/cart/checkout", middlewareChain(controller.CheckoutCartHandler)).Methods(http.MethodPost)
	router.HandleFunc("/cart/checkout/verify-payment/{reference}", middlewareChain(controller.VerifyPaymentHandler)).Methods(http.MethodGet)
	router.HandleFunc("/cart", middlewareChain(controller.GetCartHandler)).Methods(http.MethodGet)
	router.HandleFunc("/cart", middlewareChain(controller.DeleteCartHandler)).Methods(http.MethodDelete)

	
	
}