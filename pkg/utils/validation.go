package utils

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s interface{}) error {
	// Register custom validations if needed
	// validate.RegisterValidation("customtag", func(fl validator.FieldLevel) bool { ... })

	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	// Format validation errors
	var errors []string
	for _, err := range err.(validator.ValidationErrors) {
		field := strings.ToLower(err.Field())
		tag := err.Tag()
		param := err.Param()

		switch tag {
		case "required":
			errors = append(errors, field+" is required")
		case "min":
			errors = append(errors, field+" must be at least "+param+" characters")
		case "max":
			errors = append(errors, field+" must be at most "+param+" characters")
		case "email":
			errors = append(errors, field+" must be a valid email address")
		default:
			errors = append(errors, field+" is invalid")
		}
	}

	return &ValidationError{Errors: errors}
}

type ValidationError struct {
	Errors []string
}

func (e *ValidationError) Error() string {
	return strings.Join(e.Errors, ", ")
}

func init() {
	// Custom field name tags
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}
