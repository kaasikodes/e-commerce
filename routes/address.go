package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaasikodes/e-commerce-go/controllers"
	"github.com/kaasikodes/e-commerce-go/types"
)

type AddressRoutes struct {
	addressRepo  types.AddressRepository
	
}

func NewAddressRoutes( addressRepo  types.AddressRepository) *AddressRoutes {

	return &AddressRoutes{
		addressRepo: addressRepo,
	}
	
}

func (c *AddressRoutes) RegisterAddressRoutes (router *mux.Router){
	controller := controllers.NewAddressController(c.addressRepo)
	
	router.HandleFunc("/addresses", controller.GetAddressesHandler).Methods(http.MethodGet)


	
	
}