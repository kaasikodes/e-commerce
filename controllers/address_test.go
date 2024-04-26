package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/services"
	"github.com/kaasikodes/e-commerce-go/types"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type GetAddressResponse struct {
	Message    string                             `json:"message"`
	StatusCode int                                `json:"statusCode"`
	Data       types.PaginatedAddressesDataOutput `json:"data"`
}

func TestAddressController_GetAddressesHandler(t *testing.T) {
	// Set up mock database and expectations
	query := `SELECT a.ID, a.StreetAddress, a.LgaID, a.StateID, a.CountryID, l.ID, l.Name, s.ID, s.Name, c.ID, c.Name, (SELECT COUNT(*) FROM Address) AS total_addresses FROM Address a JOIN Lga l ON l.ID = a.LgaID JOIN State s ON s.ID = l.StateID JOIN Country c ON c.ID = s.CountryID WHERE a.ID > ? ORDER BY a.ID ASC LIMIT ?`
	escapedQuery := regexp.QuoteMeta(query)
	db, mock, ew := sqlmock.New()
	if ew != nil {
		t.Fatalf("an ewor '%s' was not expected when opening a stub database connection", ew)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"ID", "StreetAddress", "LgaID", "StateID", "CountryID", "ID", "LgaName", "ID", "StateName", "ID", "CountryName", "total_addresses"}).
		AddRow(1, "test address", 1, 1, 1, 1, "Lga 1", 1, "State 1", 1, "Country 1", 1).
		AddRow(2, "test address", 1, 1, 1, 1, "Lga 1", 1, "State 1", 1, "Country 1", 2)
	mock.ExpectPrepare(escapedQuery).ExpectQuery().WithArgs("", constants.DefaultPageSize).WillReturnRows(rows)

	// Create controller with mocked repository
	repo := services.NewAddressRepository(db)
	controller := NewAddressController(repo)

	// Create a request
	req := httptest.NewRequest(http.MethodGet, "/address", nil)
	if ew != nil {
		t.Fatal(ew)
	}

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Create a handler function
	controller.GetAddressesHandler(w, req)

	// Check the status code
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	//convert respose to json
	var response GetAddressResponse
	body := w.Body.Bytes()
	if err := json.Unmarshal(body, &response); err != nil {
		t.Fatalf("error parsing response body: %v", err)
	}
	fmt.Println(response, "RESPONSE")
	// Check the response body -> message, items, total ...
	if response.StatusCode != http.StatusOK {
		t.Errorf("handler returned unexpected message: got %v want %v", response.StatusCode, http.StatusOK)
	}
	expectedMsg := "Addresses retrieved successfully!"
	if response.Message != expectedMsg {
		t.Errorf("handler returned unexpected message: got %v want %v", response.Message, expectedMsg)
	}
	expectedTotal := 2
	if response.Data.Total != expectedTotal {
		t.Errorf("handler returned unexpected message: got %v want %v", response.Data.Total, expectedTotal)

	}

	// This depends on what your handler actually returns
	// You need to inspect the actual response returned by your handler
	// For now, let's just check if it's not empty
	if w.Body.Len() == 0 {
		t.Errorf("handler returned empty body")
	}
}
