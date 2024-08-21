package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gitlab.bd.com/new-argos-be/cmd/api/handlers"
	"gitlab.bd.com/new-argos-be/common"
)

type Application struct {
	logger  echo.Logger
	server  *echo.Echo
	handler handlers.Handler
}

func main() {
	e := echo.New()
	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal("Error loading the file")
	}
	db, err := common.NewDB()
	if err != nil {
		e.Logger.Fatal(err.Error())
	}

	h := handlers.Handler{
		DB: db,
	}

	app := Application{
		logger:  e.Logger,
		server:  e,
		handler: h,
	}

	app.routes(h)

	port := os.Getenv("APP_PORT")
	appAddress := fmt.Sprintf(":%s", port)
	e.Logger.Fatal(e.Start(appAddress))
}
