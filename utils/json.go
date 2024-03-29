package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kaasikodes/e-commerce-go/types"
)



func WriteJson(w http.ResponseWriter, status int, message string, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(types.ApiResponse{
		StatusCode: status,
		Message: message,
		Data: data,
	})

}

func WriteError (w http.ResponseWriter , status int, message string, errs []error) error{
	errMessages := []string{}
	for _, err := range errs {
		errMessages = append(errMessages, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(types.ApiError{
		StatusCode: status,
		Message: message,
		Errors: errMessages,
	})
	
}

func ParseJSON (r *http.Request, v interface{}) error {
	if(r.Body == nil){
		return fmt.Errorf("request body is empty")
	}
	return json.NewDecoder(r.Body).Decode(v)
}