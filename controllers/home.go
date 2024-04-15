package controllers

import (
	"net/http"

	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/utils"
)

type HomeController struct {
}

func NewHomeController() *HomeController {
	return &HomeController{
	}
}

func (c *HomeController) GetHomeBaseHandler(w http.ResponseWriter, r *http.Request)  {
	
	utils.WriteJson(w, http.StatusOK, "Welcome to e-commerce!",  map[string]string{"message": "Welcome to e-commerce!", "status": "success", "version": constants.AppVersion})
		
}

