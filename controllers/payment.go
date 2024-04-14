package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/types"
	"github.com/kaasikodes/e-commerce-go/utils"
)

type PaymentController struct {
	paymentRepo types.PaymentRepository
}

func NewPaymentController(paymentRepo types.PaymentRepository) *PaymentController {
	return &PaymentController{
		paymentRepo: paymentRepo,
	}
}


func (c *PaymentController) GetPaymentsHandler(w http.ResponseWriter, r *http.Request)  {
	user, err := utils.RetrieveUserFromRequestContext(r)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	customerId := user.Customer.ID;
	repo := c.paymentRepo
	// get query params
	pageSizeStr := r.URL.Query().Get(constants.QueryPageSize)
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil && pageSizeStr != "" {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{constants.ErrPageSizeNotValid})
		return
	}
	// get payments
	payments, err := repo.RetrievePayments( types. RetrievePaymentsInput{Pagination: types.Pagination{
		PageSize: pageSize,
	}}, customerId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, "Payments retrieved successfully!",  payments)
		
}


func (c *PaymentController) GetPaymentHandler(w http.ResponseWriter, r *http.Request)  {
	repo := c.paymentRepo
	// get query params
	id := mux.Vars(r)["id"]

	
	payment, err := repo.RetrievePayment(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, "Payment retrieved successfully!",  payment)
		
}
