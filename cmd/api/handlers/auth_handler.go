package handlers

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/labstack/echo/v4"
	"gitlab.bd.com/new-argos-be/cmd/api/requests"
	"gitlab.bd.com/new-argos-be/cmd/api/services"
	"gitlab.bd.com/new-argos-be/common"
	"gorm.io/gorm"
)

func (v *Handler) RegisterHandler(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		c.Logger().Error("Failed to read request body:", err)
		return common.SendBadRequestResponse(c, err.Error())
	}

	payload := new(requests.RegisterTerminalRequest)
	if err := json.Unmarshal(body, payload); err != nil {
		c.Logger().Error("Failed to unmarshal request body:", err)
		return common.SendBadRequestResponse(c, err.Error())
	}

	c.Logger().Infof("Received payload: %+v", payload)
	validationErrors := v.ValidateBodyRequest(c, *payload)

	if validationErrors != nil {
		return common.SendFailedValidationResponse(c, validationErrors)
	}

	userService := services.NewUserService(v.DB)

	terminalExists, err := userService.GetTerminalByIp(payload.IPv4)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return common.SendBadRequestResponse(c, "Ipv4 já cadastrado!")
	}
	print(terminalExists)
	userService.RegisterTerminal(payload)
	return common.SendSuccessResponse(c, "User registration successful.", nil)
}
