package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/models"
	"github.com/kaasikodes/e-commerce-go/types"
	"github.com/kaasikodes/e-commerce-go/utils"
)

type CategoryController struct {
	categoryRepo types.CategoryRepository
}

func NewCategoryController(categoryRepo types.CategoryRepository) *CategoryController {
	return &CategoryController{
		categoryRepo: categoryRepo,
	}
}


func (c *CategoryController) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request)  {
	repo := c.categoryRepo
	id := mux.Vars(r)["id"]

	

	// get category
	category, err := repo.DeleteCategory(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, "Category deleted successfully!",  category)
		
}
func (c *CategoryController) EditCategoryHandler(w http.ResponseWriter, r *http.Request)  {
	repo := c.categoryRepo
	id := mux.Vars(r)["id"]

	var payload models.Category
	
	if err:= utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{err})
		return
		
	}
	errParsed := utils.ValidatePayload(payload)
	if len(errParsed) > 0{
	
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, errParsed)
		return
	}
	// get category
	category, err := repo.UpdateCategory(id, models.Category{
		Name: payload.Name,
		Description: payload.Description,

	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, "Category updated successfully!",  category)
		
}

func (c *CategoryController) AddCategoryHandler(w http.ResponseWriter, r *http.Request)  {
	repo := c.categoryRepo
	var payload models.Category
	
	if err:= utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{err})
		return
		
	}
	errParsed := utils.ValidatePayload(payload)
	if len(errParsed) > 0{
	
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, errParsed)
		return
	}
	// get category
	category, err := repo.AddCategory(models.Category{
		Name: payload.Name,
		Description: payload.Description,

	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, constants.MsgCategoryRetrieved,  category)
		
}
// csv columns
const (
	name int = iota
	description
)

func (c *CategoryController) ImportMultipleCategoryHandler(w http.ResponseWriter, r *http.Request)  {
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
	categoriesToBeAdded := []types.MultipleCategoryInput{}
	// discard header from csv record
	record = record[1:]
	errParsed := []error{}

	for i, row := range record {
		category := types.MultipleCategoryInput{Name: row[name], Description: row[description]}
		if(len(utils.ValidatePayload(category)) > 0){//TODO: Find an elegant n pretty way to return the err messages
			errParsed = append(errParsed, fmt.Errorf("row %d",i+1 ))

			errParsed = append(errParsed,  utils.ValidatePayload(category)...)
		}
		categoriesToBeAdded= append(categoriesToBeAdded, category)
	}
	if len(errParsed) > 0{
	
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, errParsed)
		return
	}
	categories, err := c.categoryRepo.AddMultipleCategories(categoriesToBeAdded)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}

	utils.WriteJson(w, http.StatusOK, "Categories imported successfully!",  categories)
		
}
func (c *CategoryController) GetImportCategoryTemplateHandler(w http.ResponseWriter, r *http.Request)  {
	// TODO: Refactor to reusable fn in utilities
	
	csvData := [][]string{
		{"Name", "Description"},
		{"Groceries", "A category to hold toiletries, cereals, ...."},
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
func (c *CategoryController) GetCategoryHandler(w http.ResponseWriter, r *http.Request)  {
	repo := c.categoryRepo
	// get query params
	id := mux.Vars(r)["id"]

	// get category
	category, err := repo.RetrieveCategoryByID(id,)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, constants.MsgCategoryRetrieved,  category)
		
}


func (c *CategoryController) GetCategoriesHandler(w http.ResponseWriter, r *http.Request)  {
	user, err := utils.RetrieveUserFromRequestContext(r)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	fmt.Println("User is: ", user)
	repo := c.categoryRepo
	// get query params
	pageSizeStr := r.URL.Query().Get(constants.QueryPageSize)
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil && pageSizeStr != "" {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{constants.ErrPageSizeNotValid})
		return
	}
	// get categories
	categories, err := repo.RetrieveCategories( types.RetrievCategoriesInput{Pagination: types.Pagination{
		PageSize: pageSize,
	}})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, constants.MsgCategoriesRetrieved,  categories)
		
}