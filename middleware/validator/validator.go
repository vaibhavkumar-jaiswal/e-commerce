package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Generic validation function
func ValidateStruct(s any) error {
	err := validate.Struct(s)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			fmt.Printf("\nValidation failed for field '%s': %s\n", err.Field(), err.ActualTag())
		}
		return err
	}
	return nil
}
