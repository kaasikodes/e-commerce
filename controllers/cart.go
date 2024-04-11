package controllers

import (
	"net/http"

	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/types"
	"github.com/kaasikodes/e-commerce-go/utils"
)

type CartController struct {
	cartRepo types.CartRepository
}

func NewCartController(cartRepo types.CartRepository) *CartController {
	return &CartController{
		cartRepo: cartRepo,
	}
}


func (c *CartController) DeleteCartHandler(w http.ResponseWriter, r *http.Request)  {
	repo := c.cartRepo
	user, err := utils.RetrieveUserFromRequestContext(r)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	customerId := user.Customer.ID

	

	// delete
	err = repo.DeleteCart(customerId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, "Cart removed successfully!",  nil)
		
}


func (c *CartController) SaveCartHandler(w http.ResponseWriter, r *http.Request)  {
	repo := c.cartRepo
	user, err := utils.RetrieveUserFromRequestContext(r)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	customerId := user.Customer.ID
	_, err = repo.RetrieveCart(customerId)
	if err == nil { //delete cart if it exists
		err := repo.DeleteCart(customerId)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
			return
		}
	}
	var payload types.SaveCartInput
	
	if err:= utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{err})
		return
		
	}
	errParsed := utils.ValidatePayload(payload)
	if len(errParsed) > 0{
	
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, errParsed)
		return
	}
	// create cart
	cart, err := repo.CreateCart(payload, customerId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, "Cart saved successfully!",  cart)
		
}

func (c *CartController) GetCartHandler(w http.ResponseWriter, r *http.Request)  {
	repo := c.cartRepo
	user, err := utils.RetrieveUserFromRequestContext(r)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	customerId := user.Customer.ID

	// get cart
	cart, err := repo.RetrieveCart(customerId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Unable to retrieve customer cart for user!", []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, "Customer cart retieved successfully!",  cart)
		
}


