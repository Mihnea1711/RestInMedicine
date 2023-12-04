package validation

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

// Register custom validation tags
func RegisterCustomValidationTags() {
	validate := validator.New()
	validate.RegisterValidation("isbool", ValidateIsBool)
}

// Custom validation function for isbool tag
func ValidateIsBool(fl validator.FieldLevel) bool {
	field := fl.Field()
	kind := field.Kind()
	return kind == reflect.Bool
}
