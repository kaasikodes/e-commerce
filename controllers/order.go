package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/types"
	"github.com/kaasikodes/e-commerce-go/utils"
)

type OrderController struct {
	orderRepo types.OrderRepository
	userRepo types.UserRepository
}

func NewOrderController(orderRepo types.OrderRepository,  userRepo types.UserRepository) *OrderController {
	return &OrderController{
		orderRepo: orderRepo,
		userRepo: userRepo,
	}
}


func (c *OrderController) GetOrdersHandler(w http.ResponseWriter, r *http.Request)  {
	user, err := utils.RetrieveUserFromRequestContext(r)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	customerId := user.Customer.ID;
	repo := c.orderRepo
	// get query params
	pageSizeStr := r.URL.Query().Get(constants.QueryPageSize)
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil && pageSizeStr != "" {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{constants.ErrPageSizeNotValid})
		return
	}
	// get orders
	orders, err := repo.RetrieveOrders( types.RetrievOrdersInput{Pagination: types.Pagination{
		PageSize: pageSize,
	}}, customerId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, "Orders retrieved successfully!",  orders)
		
}

func (c *OrderController) DeleteOrderHandler(w http.ResponseWriter, r *http.Request)  {
	repo := c.orderRepo
	id := mux.Vars(r)["id"]

	

	// delete product
	 err := repo.DeleteOrder(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, "Order deleted successfully!",  nil)
		
}

func (c *OrderController) GetOrderHandler(w http.ResponseWriter, r *http.Request)  {
	repo := c.orderRepo
	// get query params
	id := mux.Vars(r)["id"]

	
	order, err := repo.RetrieveOrder(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, "Order retrieved successfully!",  order)
		
}
