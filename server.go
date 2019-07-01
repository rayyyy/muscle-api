package main

import (
	"muscle-api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello", controllers.Hello)
	e.Logger.Fatal(e.Start(":80"))
}
