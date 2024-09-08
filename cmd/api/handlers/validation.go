package handlers

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gitlab.bd.com/new-argos-be/common"
)

func (v *Handler) ValidateBodyRequest(c echo.Context, payload any) []*common.ValidationErrors {
	validate := validator.New(validator.WithRequiredStructEnabled())
	var errors []*common.ValidationErrors
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
			param := validationErr.Param()
			errMessage := keyToTitleCase + " field is " + condition
			switch condition {
			case "required":
				errMessage = keyToTitleCase + " is required"
			case "email":
				errMessage = keyToTitleCase + " must be a valid email address"
			case "min":
				if _, err := strconv.Atoi(param); err == nil {
					errMessage = fmt.Sprintf("%s must be at least %s characters", keyToTitleCase, param)
				}
			case "containsany":
				errMessage = fmt.Sprintf("%s needs to have at least %s", keyToTitleCase, param)
			}
			currentValidationError := &common.ValidationErrors{
				Error:     errMessage,
				Key:       keyToTitleCase,
				Condition: condition,
			}
			errors = append(errors, currentValidationError)
		}
	}
	return errors
}
