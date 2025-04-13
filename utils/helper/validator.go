package helper

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func validateStruct(s any) map[string]string {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	fmt.Println("Validation error:", err)
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return map[string]string{
			"error": "Validation error",
		}
	}

	fieldMap := getJSONFieldMap(s)
	errorsMap := make(map[string]string)

	for _, fieldErr := range validationErrors {
		fieldName := fieldErr.StructField()
		jsonField := fieldMap[fieldName]
		if jsonField == "" {
			jsonField = strings.ToLower(fieldName)
		}

		var msg string
		switch fieldErr.Tag() {
		case "required":
			msg = "must be required"
		case "email":
			msg = "must be a valid email address"
		case "min":
			msg = fmt.Sprintf("must be at least %s characters", fieldErr.Param())
		case "max":
			msg = fmt.Sprintf("must be at most %s characters", fieldErr.Param())
		default:
			msg = "is invalid"
		}

		errorsMap[jsonField] = msg
	}

	return errorsMap
}

func getJSONFieldMap(s any) map[string]string {
	fieldMap := make(map[string]string)

	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" && jsonTag != "-" {
			jsonName := strings.Split(jsonTag, ",")[0]
			fieldMap[field.Name] = jsonName
		}
	}

	return fieldMap
}
