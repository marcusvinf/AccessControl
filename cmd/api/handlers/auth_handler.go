package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func (v *Handler) RegisterHandler(c echo.Context) error {
	fmt.Println("test")
}
