package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/types"
	"github.com/kaasikodes/e-commerce-go/utils"
)

type CartController struct {
	cartRepo types.CartRepository
	orderRepo types.OrderRepository
	paymentRepo types.PaymentRepository
	addressRepo types.AddressRepository
}

func NewCartController(cartRepo types.CartRepository, orderRepo types.OrderRepository, paymentRepo types.PaymentRepository, addressRepo types.AddressRepository) *CartController {
	return &CartController{
		cartRepo: cartRepo,
		orderRepo: orderRepo,
		paymentRepo: paymentRepo,
		addressRepo: addressRepo,
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


func (c *CartController) VerifyPaymentHandler(w http.ResponseWriter, r *http.Request)  {
	
	cartRepo := c.cartRepo
	paymentRepo := c.paymentRepo
	reference := mux.Vars(r)["reference"]
	verfifyResponse, err := cartRepo.VerifyPayment(reference)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error while verifying payment!", []error{err})
		return
	}
	// if payment is successful
	// update order payment in db
	// send mail teling user that payment has been made
	var finalRes = make(map[string]string)
	switch verfifyResponse.Data.Status {
		case constants.PaystackTransactionStatuses.Success:
			// update the payment status of order in database
			err = paymentRepo.UpdatePayment(types.UpdatePaymentInput{
				Paid: true,
				PaidAt: verfifyResponse.Data.PaidAt,
			}, reference)
			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, "Error while verifying payment!", []error{err})
				return
			}
			finalRes["paid"] = "true"
		case constants.PaystackTransactionStatuses.Abandoned:
			finalRes["paid"] = "false"
		default:
			finalRes["paid"] = "false"
	}
	

	
	utils.WriteJson(w, http.StatusOK, "Payment verified!",  finalRes)
		
}
func (c *CartController) CheckoutCartHandler(w http.ResponseWriter, r *http.Request)  {
	// TODO: Consider making a db transaction for this, since multple tables are involved
	cartRepo := c.cartRepo
	orderRepo := c.orderRepo
	paymentRepo := c.paymentRepo
	addressRepo := c.addressRepo

	var payload types.CartCheckoutInput
	
	if err:= utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{err})
		return
		
	}
	errParsed := utils.ValidatePayload(payload)
	if len(errParsed) > 0{
	
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, errParsed)
		return
	}
	user, err := utils.RetrieveUserFromRequestContext(r)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	customerId := user.Customer.ID
	userEmail := user.Email
	_, err = cartRepo.RetrieveCart(customerId)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Unable to retrieve customer cart for user!", []error{err})
		return
	}
	virtualOrder, err := cartRepo.CheckoutCart(customerId, userEmail)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Unable to retrieve virtual order for user!", []error{err})
		return
	}
	
	// create order in db
	var createOrderInput types.CreateOrderInput
	createOrderInput.TotalAmount = virtualOrder.TotalAmount
	createOrderInput.OrderItems = []types.OrderItemInput{}
	for _, item := range virtualOrder.Items {
		createOrderInput.OrderItems = append(createOrderInput.OrderItems, types.OrderItemInput{
			ProductId: item.ProductID,
			TotalPrice: item.TotalPrice,
			Quantity: item.Quantity,
		})
	}
	virtualOrder.DeliveryAddressID, err = addressRepo.CreateAddress(payload.DeliveryAddress)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error while creating address for cart!", []error{err})
		return
	}
	orderId, err := orderRepo.CreateOrder(createOrderInput, customerId, virtualOrder.DeliveryAddressID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error while creating order for cart!", []error{err})
		return
	}

	// create payment in db
	err = paymentRepo.CreatePayment(types.CreatePaymentInput{
		Reference: virtualOrder.Payment.ID,
		Amount: virtualOrder.TotalAmount,
		

	}, orderId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error while saving payment for cart!", []error{err})
		return
	}
	// delete cart
	err = cartRepo.DeleteCart(customerId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}

	
	utils.WriteJson(w, http.StatusOK, "Succesful cart checkout!",  virtualOrder)
		
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


