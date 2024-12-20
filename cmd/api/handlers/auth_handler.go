package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
	"gitlab.bd.com/new-argos-be/cmd/api/requests"
	"gitlab.bd.com/new-argos-be/cmd/api/services"
	"gitlab.bd.com/new-argos-be/common"
	"gorm.io/gorm"
)

func (v *Handler) RegisterTerminalHandler(c echo.Context) error {
	userService := services.NewUserService(v.DB)
	_, err := userService.GetAllLocations()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return common.SendBadRequestResponse(c, "Nenhuma localidade cadastradada!")
	}
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		c.Logger().Error("Falha ao ler corpo da requisicao:", err)
		return common.SendBadRequestResponse(c, err.Error())
	}

	payload := new(requests.RegisterTerminalRequest)
	if err := json.Unmarshal(body, payload); err != nil {
		c.Logger().Error("Falha ao deserializar o corpo da requisicao:", err)
		return common.SendBadRequestResponse(c, err.Error())
	}

	c.Logger().Infof("Received payload: %+v", payload)
	validationErrors := v.ValidateBodyRequest(c, *payload)

	if validationErrors != nil {
		return common.SendFailedValidationResponse(c, validationErrors)
	}

	terminalExists, err := userService.GetTerminalByIp(payload.IPv4)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return common.SendBadRequestResponse(c, fmt.Sprintf("IPv4 já registrado, terminal de Nome %s", terminalExists.Name))
	}
	localId, err := userService.GetLocalByName(payload.LocalName)
	if err != nil {
		return common.SendBadRequestResponse(c, err.Error())
	}
	userService.RegisterTerminal(payload, localId.LocalID)
	return common.SendSuccessResponse(c, "Terminal registrado com sucesso", nil)
}

func (v *Handler) RegisterLocalHandler(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		c.Logger().Error("Falha ao ler corpo da requisicao:", err)
		return common.SendBadRequestResponse(c, err.Error())
	}

	payload := new(requests.RegisterLocalRequest)
	if err := json.Unmarshal(body, payload); err != nil {
		c.Logger().Error("Falha ao deserializar o corpo da requisicao:", err)
		return common.SendBadRequestResponse(c, err.Error())
	}

	c.Logger().Infof("Payload recebido: %+v", payload)
	validationErrors := v.ValidateBodyRequest(c, *payload)

	if validationErrors != nil {
		return common.SendFailedValidationResponse(c, validationErrors)
	}

	userService := services.NewUserService(v.DB)

	locationExists, err := userService.GetLocalByName(payload.Name)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return common.SendBadRequestResponse(c, "Local "+locationExists.Name+" já cadastrado!")
	}

	registeredLocal, err := userService.RegisterLocal(payload)

	if err != nil {
		return common.SendInternalServerErrorResponse(c, err.Error())
	}
	return common.SendSuccessResponse(c, "Local registrado com sucesso", registeredLocal)
}

func (v *Handler) LoginHandler(c echo.Context) error {
	userService := services.NewUserService(v.DB)
	payload := new(requests.LoginRequest)
	if err := (&echo.DefaultBinder{}).BindBody(c, payload); err != nil {
		c.Logger().Error(err)
		return common.SendBadRequestResponse(c, err.Error())
	}

	validationErrors := v.ValidateBodyRequest(c, *payload)
	if validationErrors != nil {
		return common.SendFailedValidationResponse(c, validationErrors)
	}
	userRetrieved, err := userService.GetUserByUsername(payload.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return common.SendBadRequestResponse(c, "Usuário ja cadastrado")
	}

	if !common.CheckPasswordHash(payload.Password, userRetrieved.Password) {
		return common.SendBadRequestResponse(c, "Credenciais invalidas!")
	}

	accessToken, refreshToken, err := common.GenerateJWT(*userRetrieved)
	if err != nil {
		return common.SendInternalServerErrorResponse(c, err.Error())
	}
	return common.SendSuccessResponse(c, "Usuário logado!", map[string]any{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user":          userRetrieved,
	})
}
