package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaasikodes/e-commerce-go/controllers"
	middleware "github.com/kaasikodes/e-commerce-go/middlware"
	"github.com/kaasikodes/e-commerce-go/types"
)

type PaymentRoutes struct {
	paymentRepo types.PaymentRepository
	userRepo types.UserRepository
}

func NewPaymentRoutes(  paymentRepo types.PaymentRepository,  userRepo types.UserRepository) *PaymentRoutes {
	return &PaymentRoutes{
		paymentRepo: paymentRepo,
		userRepo: userRepo,
	}
}

func (c *PaymentRoutes) RegisterPaymentRoutes (router *mux.Router){
	controller := controllers.NewPaymentController(c.paymentRepo)
	middlewareChain := middleware.MiddlewareChain(middleware.RequireAuthMiddleware(c.userRepo))
	
	router.HandleFunc("/payments/{id}", middlewareChain(controller.GetPaymentHandler)).Methods(http.MethodGet)
	router.HandleFunc("/payments", middlewareChain(controller.GetPaymentsHandler)).Methods(http.MethodGet)


	
	
}