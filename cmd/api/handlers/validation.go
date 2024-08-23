package handlers

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ValidationErrors struct {
	Error     string `json:"error"`
	Key       string `json:"key"`
	Condition string `json:"condition"`
}

func (v *Handler) ValidateBodyRequest(c echo.Context, payload any) []*ValidationErrors {
	validate := validator.New(validator.WithRequiredStructEnabled())
	var errors []*ValidationErrors
	err := validate.Struct(payload)
	validationErrors, ok := err.(validator.ValidationErrors)
	if ok {
		reflected := reflect.ValueOf(payload)
		for _, validationErr := range validationErrors {
			field, _ := reflected.Type().FieldByName(validationErr.StructField())
			key := field.Tag.Get("json")

			if key == "" {
				key = strings.ToLower(validationErr.StructField())
			}
			condition := validationErr.Tag()
			keyToTitleCase := strings.Replace(key, "_", " ", -1)
			errMessage := keyToTitleCase + " field is " + condition
			currentValidationError := &ValidationErrors{
				Error:     errMessage,
				Key:       keyToTitleCase,
				Condition: condition,
			}
			errors = append(errors, currentValidationError)
		}
		// return errors
	}
	// currentValidationError := &ValidationErrors{
	// 	Error:     "",
	// 	Key:       "",
	// 	Condition: "",
	// }
	// errors = append(errors, currentValidationError)
	return errors
}
