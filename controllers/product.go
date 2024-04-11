package controllers

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/models"
	"github.com/kaasikodes/e-commerce-go/types"
	"github.com/kaasikodes/e-commerce-go/utils"
)

type ProductController struct {
	productRepo types.ProductRepository
	categoryRepo types.CategoryRepository
}

func NewProductController(productRepo types.ProductRepository , categoryRepo types.CategoryRepository) *ProductController {
	return &ProductController{
		productRepo: productRepo,
		categoryRepo: categoryRepo,
	}
}

func (c *ProductController) AddProductHandler(w http.ResponseWriter, r *http.Request)  {
	repo := c.productRepo
	var payload types.AddProductInput
	
	if err:= utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{err})
		return
		
	}
	errParsed := utils.ValidatePayload(payload)
	if len(errParsed) > 0{
	
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, errParsed)
		return
	}
	user, err := utils.RetrieveUserFromRequestContext(r);
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	fmt.Println(user, "user in product")
	sellerId := user.Seller.ID;
	fmt.Println(sellerId, "seller id in product")
	// add product
	product, err := repo.AddProduct(payload, sellerId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, "Product added successfully!",  product)
		
}

func (c *ProductController) DeleteProductHandler(w http.ResponseWriter, r *http.Request)  {
	repo := c.productRepo
	id := mux.Vars(r)["id"]

	

	// delete product
	product, err := repo.DeleteProduct(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, "Product deleted successfully!",  product)
		
}

func (c *ProductController) EditProductHandler(w http.ResponseWriter, r *http.Request)  {
	repo := c.productRepo
	id := mux.Vars(r)["id"]

	var payload types.AddProductInput
	
	if err:= utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{err})
		return
		
	}
	errParsed := utils.ValidatePayload(payload)
	if len(errParsed) > 0{
	
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, errParsed)
		return
	}
	
	product, err := repo.UpdateProduct(id, payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, "Product updated successfully!",  product)
		
}

func (c *ProductController) GetProductHandler(w http.ResponseWriter, r *http.Request)  {
	repo := c.productRepo
	// get query params
	id := mux.Vars(r)["id"]

	
	product, err := repo.RetrieveProductByID(id,)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, "Product retrieved successfully!",  product)
		
}

func (c *ProductController) GetProductsHandler(w http.ResponseWriter, r *http.Request)  {
	user, err := utils.RetrieveUserFromRequestContext(r)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	sellerId := user.Seller.ID;
	repo := c.productRepo
	// get query params
	pageSizeStr := r.URL.Query().Get(constants.QueryPageSize)
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil && pageSizeStr != "" {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{constants.ErrPageSizeNotValid})
		return
	}
	// get products
	products, err := repo.RetrieveProducts( types.RetrievProductsInput{Pagination: types.Pagination{
		PageSize: pageSize,
	}}, sellerId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, "Products retrieved successfully!",  products)
		
}
const (
	productName int = iota 
	productDescription
	productPrice
	productQuantity
	productCategory //TODO: Add a unique restraint to the category name in db, 
)

func getCategoryIdFromName (categories []models.Category ,name string) (string, error) {
	for _, category := range categories {
		if category.Name == name {
			return category.ID, nil
		}
	}
	return "", errors.New("category not found")
	
}
func (c *ProductController) ImportMultipleProductHandler(w http.ResponseWriter, r *http.Request)  {
	// TODO: Refactor to reusable fn in utilities
	err := r.ParseMultipartForm(10 * 10 *1024);
	if err != nil{
		utils.WriteError(w, http.StatusBadRequest, "Invalid form data", []error{err})
		return
	}

	file, _, err := r.FormFile("csvFile")
	if err !=nil {
		utils.WriteError(w, http.StatusBadRequest, "csvFile absent", []error{err})
		return
	}
	// TODO: Validate that file type is csv
	
	defer file.Close()	
	// create a csv reader
	reader := csv.NewReader(file)
	record, err := reader.ReadAll()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	productsToBeAdded := []types.MultipleProductInput{}
	// discard header from csv record
	record = record[1:]
	errParsed := []error{}
	categories, err := c.categoryRepo.RetrieveCategories(types.RetrievCategoriesInput{
		Pagination: types.Pagination{
			PageSize: 100,
		},
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	for i, row := range record {
		// convert price to number
		price, err := strconv.ParseInt(row[productPrice], 10, 64)
		if err != nil {
			errParsed = append(errParsed, fmt.Errorf("row %d price could not be parsed to an integer",i+1 ))
			errParsed = append(errParsed, err)
		}
		// convert quantity to number
		quantity, err := strconv.ParseInt(row[productQuantity], 10, 64)
		if err != nil {
			errParsed = append(errParsed, fmt.Errorf("row %d quantity could not be parsed to an integer",i+1 ))
			errParsed = append(errParsed, err)
		}
		// get category id from name
		categoryId, err := getCategoryIdFromName(categories.Data, row[productCategory])
		if err != nil {
			errParsed = append(errParsed, fmt.Errorf("row %d %s category cannot be found",i+1, row[productCategory] ))
			errParsed = append(errParsed, err)
			
		}
		product := types.MultipleProductInput{
			Name: row[productName],
			Description: row[productDescription],
			Price: int(price),
			Quantity: int(quantity),
			CategoryID: categoryId,
		}
		if(len(utils.ValidatePayload(product)) > 0){//TODO: Find an elegant n pretty way to return the err messages
			errParsed = append(errParsed, fmt.Errorf("row %d",i+1 ))

			errParsed = append(errParsed,  utils.ValidatePayload(product)...)
		}
		productsToBeAdded= append(productsToBeAdded, product)
	}
	if len(errParsed) > 0{
	
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, errParsed)
		return
	}

	user, err := utils.RetrieveUserFromRequestContext(r)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	sellerId := user.Seller.ID;
	products, err := c.productRepo.AddMultipleProducts(productsToBeAdded, sellerId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}

	utils.WriteJson(w, http.StatusOK, "Products imported successfully!",  products)
		
}
func (c *ProductController) GetImportProductTemplateHandler(w http.ResponseWriter, r *http.Request)  {
	
	csvData := [][]string{
		{"Name", "Description", "Price", "Quantity", "Category"},
		{"Groceries", "A category to hold toiletries, cereals, ....", "100", "10", "Luxury"},
	}

	// set headers for browser to download
	w.Header().Set("Content-type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=template.csv")
	// create a csv writer using our http response writer as our io.writer
	wr := csv.NewWriter(w)

	if err:=wr.WriteAll(csvData); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
		
}
