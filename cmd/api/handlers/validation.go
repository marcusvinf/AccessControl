package handlers

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ValidationErrors struct {
	Error     string `json:"error"`
	Key       string `json:"key"`
	Condition string `json:"condition"`
}

func (v *Handler) ValidaBodyRequest(c echo.Context, payload any) []*ValidationErrors {
	validate := validator.New(validator.WithRequiredStructEnabled())
	var errors []*ValidationErrors
	err := validate.Struct(payload)
	validationErrors, ok := err.(validator.ValidationErrors)
	if ok {
		reflected := reflect.ValueOf(payload)
		for _, validationErr := range validationErrors {
			fmt.Println(validationErr.Field())
			currentValidationError := &ValidationErrors{
				Error:     "e",
				Key:       "k",
				Condition: "required",
			}
			errors = append(errors, currentValidationError)
		}
	}
	return errors
}
