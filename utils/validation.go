package utils

import "github.com/go-playground/validator/v10"

var Validate = validator.New(validator.WithRequiredStructEnabled())

func ValidatePayload (payload interface{}) []error {
	errParsed := []error{}
	if err:= Validate.Struct(payload); err != nil {
		errors :=err.(validator.ValidationErrors)
		for _, e := range errors {
			errParsed = append(errParsed, e)
			
		}

	}
	return errParsed
}