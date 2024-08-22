package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.bd.com/new-argos-be/cmd/api/handlers"
	middlewares "gitlab.bd.com/new-argos-be/cmd/api/middlware"
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
		e.Logger.Fatal(err.Error())
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
	e.Use(middleware.Logger())
	e.Use(middlewares.CustomMiddleware)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:*", "http://127.0.0.1:*"}, AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept}}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{TokenLookup: "header:X-XSRF-TOKEN"}))
	app.routes(h)

	port := os.Getenv("APP_PORT")
	appAddress := fmt.Sprintf(":%s", port)
	e.Logger.Fatal(e.Start(appAddress))
}
