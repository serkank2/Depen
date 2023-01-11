package validation

import (
	"depen/model"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ---------------------Error Type---------------------
type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

// ---------------------Error Type---------------------

// ---------------------Register validation---------------------

func ValidateStruct(user model.RegisterUser) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

// ---------------------Register validation---------------------
