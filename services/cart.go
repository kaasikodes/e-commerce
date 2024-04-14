package services

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/models"
	"github.com/kaasikodes/e-commerce-go/types"
	"github.com/kaasikodes/e-commerce-go/utils"
)

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}
func (c *CartRepository) VerifyPayment(reference string) (types.VerifyPaystackTransactionResponse, error){
	var verfifyResponse types.VerifyPaystackTransactionResponse
	apiUrl := fmt.Sprintf("https://api.paystack.co/transaction/verify/%s", reference)
	ctx, cancel := context.WithTimeout(context.Background(),constants.DefaultContextTimeOut) // ensure the request does not time out
	defer cancel()
	// create new http request
	req, err := http.NewRequestWithContext(ctx, "GET", apiUrl, nil)
	if err != nil {
		return verfifyResponse,err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", constants.PaystackSecretKey))
	req.Header.Set("Content-Type", "application/json")
	// send request
	client := &http.Client{}
	res,err := client.Do(req)
	if err != nil {
		return verfifyResponse, err
	}
	resBody, error := io.ReadAll(res.Body)

    if error != nil {
        fmt.Println(error)
    }
	formattedData := formatJSON(resBody)
    fmt.Println("Status: ", res.Status)
    fmt.Println("Response body: ", formattedData)
	json.Unmarshal(resBody, &verfifyResponse)

	defer res.Body.Close()
	
	return verfifyResponse,nil
}
func (c *CartRepository) makePayment(email string, amount float64, reference string) (error){
	apiUrl := "https://api.paystack.co/transaction/initialize"
	// convert amount to kobo
	amount = amount * 100
	// convert amount to string
	amountStr:= strconv.FormatFloat(amount, 'f', -1, 64)
	fmt.Println(amountStr, "AMOUNT as string")
	fmt.Println(reference, "Refernece")
	payload := map[string]string {
		"email": email,
		"amount": amountStr,
		"reference": reference,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(),constants.DefaultContextTimeOut) // ensure the request does not time out
	defer cancel()
	// create new http request
	req, err := http.NewRequestWithContext(ctx, "POST", apiUrl, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", constants.PaystackSecretKey))
	req.Header.Set("Content-Type", "application/json")
	// send request
	client := &http.Client{}
	res,err := client.Do(req)
	if err != nil {
		return err
	}
	resBody, error := io.ReadAll(res.Body)

    if error != nil {
        fmt.Println(error)
    }
	formattedData := formatJSON(resBody)
    fmt.Println("Status: ", res.Status)
    fmt.Println("Response body: ", formattedData)
	defer res.Body.Close()
	
	return nil
}
// function to format JSON data
func formatJSON(data []byte) string {
    var out bytes.Buffer
    err := json.Indent(&out, data, "", "  ")

    if err != nil {
        fmt.Println(err)
    }

    d := out.Bytes()
    return string(d)
}
func (c *CartRepository) CheckoutCart(customerId string, userEmail string) (models.Order, error){
	
	var totalPrice float64
	orderItems := []models.OrderItem{}
	order := models.Order{}
	payment := models.Payment{}
	order.CustomerID = customerId
	orderId, _:= utils.GenerateRandomID(10)
	order.ID = orderId
	// retrieve the cart
	cart, err := c.RetrieveCart(customerId);
	if err != nil {
		return order, err
	}
	// calculate the total price
	for _, item := range cart.Items {
		itemPrice := float64(item.Product.Price) * float64(item.Quantity)
		orderItem := models.OrderItem{}
		orderItemId, _:= utils.GenerateRandomID(10)
		orderItem.ID = orderItemId
		orderItem.OrderID = orderId
		orderItem.ProductID = item.ProductID
		orderItem.Quantity = item.Quantity
		orderItem.TotalPrice = itemPrice

		totalPrice += itemPrice
		orderItems = append(orderItems, orderItem)

	}
	order.TotalAmount = totalPrice
	order.Items = orderItems
	payment.Amount = totalPrice
	payment.ID, _ = utils.GenerateRandomID(10)
	payment.OrderID = orderId
	payment.Paid = false
	// payment.Method = "cash" //this should be in constant
	order.Payment = payment
	// initiate payment
	c.makePayment(userEmail, order.TotalAmount, payment.ID)
	// wait for payment to be made (use a go routine to poll) (also create a payment repo to handle payments)
	// once made create order for user, reduce product quantity, and remove cart
	// return the order
	return order, nil
} 
func (c *CartRepository) retrieveCartItems(cartId string) ([]models.CartItem, error) {
	db := c.db
	// prepare query
	query := `
    SELECT c.*, p.*
    FROM CartItem c
    JOIN Product p ON p.ID = c.ProductID
    WHERE c.CartID = ?
	`

	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	items := []models.CartItem{}
	stmt, err := db.PrepareContext(ctx, query )
	if err !=nil {
		return items, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	rows, err := stmt.QueryContext(ctx, cartId)
	if err !=nil {
		return items, err
	}
	defer rows.Close()
	for rows.Next() {
		item := models.CartItem{}
		item.Product = models.Product{}
		err = rows.Scan(&item.ID, &item.ProductID, &item.CartID, &item.Quantity,  &item.CreatedAt, &item.UpdatedAt, &item.Product.ID, &item.Product.Name, &item.Product.Description, &item.Product.Price,&item.Product.Quantity, &item.Product.CategoryID,  &item.Product.SellerID, &item.Product.CreatedAt, &item.Product.UpdatedAt)
		if err !=nil {
			return items, err
		}
		items = append(items, item)
	}
	


	return items, nil

}
func (c *CartRepository) createCartItems(input types.SaveCartInput, cartId string) ([]models.CartItem, error) {
	db := c.db
	items := input.Items
	cartItems := []models.CartItem{}
	
	// prepare query
	query := `INSERT INTO CartItem (ID, ProductID, Quantity, CartID) VALUES`
	var inserts []string
	var params []interface{}
	for _, data := range items {
		cartItem := models.CartItem{}
		inserts = append(inserts, "(?, ?, ?, ?)")
		id, _:= utils.GenerateRandomID(10)
		params = append(params, id, data.ProductID, data.Quantity, cartId)
		cartItem.CartID = cartId
		cartItem.ProductID = data.ProductID
		cartItem.Quantity = data.Quantity;
		cartItem.Product = models.Product{}
		cartItems = append(cartItems, cartItem)

	}
	queryVals := strings.Join(inserts, ",")
	query = query + queryVals
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return cartItems, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	res, err := stmt.ExecContext(ctx, params... )
	if err !=nil {
		return cartItems, err
	}
	_, err = res.RowsAffected()
	if err !=nil {
		return cartItems, err
	}
	// ensure input have
	return cartItems, nil

}
func (c *CartRepository) CreateCart(input types.SaveCartInput, customerId string) (models.Cart, error) {
	db := c.db
	// prepare query
	query := `INSERT INTO Cart (ID, CustomerID) VALUES (?, ?)`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	cart := models.Cart{}
	if err != nil {
		return cart, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	id, _ := utils.GenerateRandomID(10)
	_, err = stmt.ExecContext(ctx, id,  customerId)
	if err != nil {
		return  cart, err
	}

	cart.ID = id
	cart.CustomerID = customerId

	cartItems, err := c.createCartItems(input, id)
	if err != nil {
		return  cart, err
	}
	cart.Items = cartItems
	return  cart, nil

}
func (c *CartRepository) RetrieveCart(customerId string) (models.Cart, error){
	db := c.db
	// prepare query
	query := `SELECT * FROM Cart WHERE CustomerID = ?`
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	cart := models.Cart{}
	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return cart, err
	}
	defer stmt.Close() //close the statement after use
	// execute the statement
	
	 err = stmt.QueryRowContext(ctx,  customerId).Scan(&cart.ID, &cart.CustomerID, &cart.CreatedAt, &cart.UpdatedAt)
	 if err !=nil {
		return cart, err
	}

	items, err := c.retrieveCartItems(cart.ID)
	if err !=nil {
		return cart, err
	}
	cart.Items = items
	
	return cart, nil

}

// Delete Cart
func (c *CartRepository)DeleteCart(customerId string) (error){
	db := c.db
	// prepare query
	query := "DELETE FROM Cart WHERE CustomerID = ?"
	// create a context as a responsible developer (to handle network error) that does not wish to waste time when something doesb't work
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultContextTimeOut)
	defer cancel()
	// prepare the statement
	stmt, err := db.PrepareContext(ctx, query)
	if err !=nil {
		return err
	}
	
	defer stmt.Close() //close the statement after use
	_, err = c.RetrieveCart(customerId)
	if err !=nil {
		return err
	}
	// execute the statement
	_, err = stmt.ExecContext(ctx, customerId)
	if err !=nil {
		return  err
	}
	
	
	return  nil
}

