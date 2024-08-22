package main

import (
	"gitlab.bd.com/new-argos-be/cmd/api/handlers"
)

func (a *Application) routes(handler handlers.Handler) {
	a.server.GET("/", handler.HealthCheck)
	a.server.POST("/register-terminal", handler.RegisterHandler)
}
