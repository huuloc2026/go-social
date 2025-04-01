package utils

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	// Register custom validators if needed
}

func ValidateRequest(data interface{}) error {
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	errorMessages := make([]string, len(validationErrors))

	for i, err := range validationErrors {
		field, _ := reflect.TypeOf(data).Elem().FieldByName(err.Field())
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = err.Field()
		}

		switch err.Tag() {
		case "required":
			errorMessages[i] = jsonTag + " is required"
		case "email":
			errorMessages[i] = jsonTag + " must be a valid email"
		case "min":
			errorMessages[i] = jsonTag + " must be at least " + err.Param() + " characters"
		case "max":
			errorMessages[i] = jsonTag + " must be at most " + err.Param() + " characters"
		default:
			errorMessages[i] = jsonTag + " is invalid"
		}
	}

	return errors.New(strings.Join(errorMessages, ", "))
}
