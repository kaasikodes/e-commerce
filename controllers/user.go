package controllers

import (
	"net/http"
	"strconv"

	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/types"
	"github.com/kaasikodes/e-commerce-go/utils"
)

type UserController struct {
	userRepo types.UserRepository
}

func NewUserController(userRepo types.UserRepository) *UserController {
	return &UserController{
		userRepo: userRepo,
	}
}

func (c *UserController) GetSellersHandler(w http.ResponseWriter, r *http.Request) {
	
	repo := c.userRepo
	// get query params
	pageSizeStr := r.URL.Query().Get(constants.QueryPageSize)
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil && pageSizeStr != "" {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{constants.ErrPageSizeNotValid})
		return
	}
	// get users
	data, err := repo.RetrieveSellers(types.RetrievUsersInput{Pagination: types.Pagination{
		PageSize: pageSize,
	}})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, "Sellers retrieved successfully!", data)

}
func (c *UserController) GetCustomersHandler(w http.ResponseWriter, r *http.Request) {
	
	repo := c.userRepo
	// get query params
	pageSizeStr := r.URL.Query().Get(constants.QueryPageSize)
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil && pageSizeStr != "" {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{constants.ErrPageSizeNotValid})
		return
	}
	// get users
	data, err := repo.RetrieveCustomers(types.RetrievUsersInput{Pagination: types.Pagination{
		PageSize: pageSize,
	}})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, "Customers retrieved successfully!", data)

}
func (c *UserController) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	
	repo := c.userRepo
	// get query params
	pageSizeStr := r.URL.Query().Get(constants.QueryPageSize)
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil && pageSizeStr != "" {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{constants.ErrPageSizeNotValid})
		return
	}
	// get users
	users, err := repo.RetrieveUsers(types.RetrievUsersInput{Pagination: types.Pagination{
		PageSize: pageSize,
	}})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, "Users retrieved successfully!", users)

}