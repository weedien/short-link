package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"shortlink/internal/base/errno"
	"strings"
)

type (
	errorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	XValidator struct {
		validator *validator.Validate
	}
)

// This is the validator instance
// for more information see: https://github.com/go-playground/validator
var validate = validator.New()

var myValidator = &XValidator{
	validator: validate,
}

func Get() *XValidator {
	return myValidator
}

func (v XValidator) Validate(data interface{}) error {
	var validationErrors []errorResponse

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem errorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	if len(validationErrors) > 0 && validationErrors[0].Error {
		errMsgs := make([]string, 0)
		for _, err := range validationErrors {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}
		return errno.NewRequestError(strings.Join(errMsgs, " and "))
	}
	return nil
}
