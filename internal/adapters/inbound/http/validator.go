package http

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	return &Validator{
		validator: validate,
	}
}

func (cv *Validator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			var errMessages []string
			for _, e := range errs {
				fieldName := e.Field()
				if e.Tag() == "required" {
					errMessages = append(errMessages, fieldName+" is required")
				} else if e.Tag() == "min" {
					errMessages = append(errMessages, fieldName+" must be greater than "+e.Param())
				} else if e.Tag() == "max" {
					errMessages = append(errMessages, fieldName+" must be less than "+e.Param())
				} else {
					errMessages = append(errMessages, fieldName+" is invalid")
				}
			}

			return errors.New(
				strings.Join(errMessages, ", "),
			)
		}
	}
	return nil
}
