package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type healthCheck struct {
	Health bool `json:"health"`
}

func (v *Handler) HealthCheck(c echo.Context) error {
	healthcheckStruct := healthCheck{
		Health: true,
	}
	return c.JSON(http.StatusOK, healthcheckStruct)
}
