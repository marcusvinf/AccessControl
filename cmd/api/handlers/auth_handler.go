package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.bd.com/new-argos-be/cmd/api/requests"
)

func (v *Handler) RegisterHandler(c echo.Context) error {
	payload := new(requests.RegisterPersonRequest)
	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	c.Logger().Info(payload)
	return c.String(http.StatusAccepted, "good request")

}
