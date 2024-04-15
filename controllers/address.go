package controllers

import (
	"net/http"
	"strconv"

	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/types"
	"github.com/kaasikodes/e-commerce-go/utils"
)

type AddressController struct {
	addressRepo types.AddressRepository
}

func NewAddressController(addressRepo types.AddressRepository ) *AddressController {
	return &AddressController{
		addressRepo: addressRepo,
	}
}

func (c *AddressController) GetAddressesHandler(w http.ResponseWriter, r *http.Request)  {

	repo := c.addressRepo
	// get query params
	pageSizeStr := r.URL.Query().Get(constants.QueryPageSize)
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil && pageSizeStr != "" {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{constants.ErrPageSizeNotValid})
		return
	}
	// get addresses
	addresses, err := repo.RetrieveAddresses( types.RetrievAddressesInput{Pagination: types.Pagination{
		PageSize: pageSize,
	}})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, "Addresses retrieved successfully!",  addresses)
		
}