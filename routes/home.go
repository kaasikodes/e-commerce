package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaasikodes/e-commerce-go/controllers"
)

type HomeRoutes struct {
	
}

func NewHomeRoutes( ) *HomeRoutes {

	return &HomeRoutes{}
	
}

func (c *HomeRoutes) RegisterHomeRoutes (router *mux.Router){
	controller := controllers.NewHomeController()
	
	router.HandleFunc("/", controller.GetHomeBaseHandler).Methods(http.MethodGet)


	
	
}