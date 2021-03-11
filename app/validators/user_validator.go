package validators

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// BookValidator func for create a new validator for expected fields,
// register function to get tag name from `json` tags.
func BookValidator() *validator.Validate {
	// Create a new validator.
	v := validator.New()

	// Get tag name from `json`.
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		// Define name of field.
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		// Processing name value.
		if name == "-" {
			return ""
		}

		return name
	})

	// Validator for user ID (UUID).
	_ = v.RegisterValidation("id", func(fl validator.FieldLevel) bool {
		// Define field as string.
		field := fl.Field().String()

		// Return true, if UUID is not valid.
		if _, err := uuid.Parse(field); err != nil {
			return true
		}

		return false
	})

	return v
}
