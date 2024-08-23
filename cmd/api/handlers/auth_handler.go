package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.bd.com/new-argos-be/cmd/api/requests"
)

func (v *Handler) RegisterHandler(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		c.Logger().Error("Failed to read request body:", err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	payload := new(requests.RegisterPersonRequest)
	if err := json.Unmarshal(body, payload); err != nil {
		c.Logger().Error("Failed to unmarshal request body:", err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	c.Logger().Infof("Received payload: %+v", payload)
	validationErrors := v.ValidateBodyRequest(c, *payload)

	if validationErrors != nil {
		return c.JSON(http.StatusOK, validationErrors)
	}

	return c.JSON(http.StatusOK, validationErrors)
}
