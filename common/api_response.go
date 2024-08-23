package common

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ValidationErrors struct {
	Error     string `json:"error"`
	Key       string `json:"key"`
	Condition string `json:"condition"`
}

type JSONSuccessResponse struct {
	Success bool   `json:"sucess"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type JSONFailedValidationResult struct {
	Success bool                `json:"sucess"`
	Message string              `json:"message"`
	Errors  []*ValidationErrors `json:"errors"`
}

type JSONErrorResponse struct {
	Success bool   `json:"sucess"`
	Message string `json:"message"`
}

func SendSuccessResponse(c echo.Context, message string, data any) error {
	return c.JSON(http.StatusOK, JSONSuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})

}

func SendFailedValidationResponse(c echo.Context, errors []*ValidationErrors) error {
	return c.JSON(http.StatusUnprocessableEntity, JSONFailedValidationResult{
		Success: false,
		Errors:  errors,
		Message: "Validation failed",
	})
}

func SendErrorResponse(c echo.Context, message string, statusCode int) error {
	return c.JSON(statusCode, JSONErrorResponse{
		Success: false,
		Message: message,
	})
}
func SendBadRequestResponse(c echo.Context, message string) error {
	return SendErrorResponse(c, message, http.StatusBadRequest)
}

func SendNotFoundResponse(c echo.Context, message string) error {
	return SendErrorResponse(c, message, http.StatusNotFound)
}

func SendInternalServerErrorResponse(c echo.Context, message string) error {
	return SendErrorResponse(c, message, http.StatusInternalServerError)
}
