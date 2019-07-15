package main

import (
	"log"
	"muscle-api/controllers"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/sign-in", controllers.SignIn)
	e.GET("/hello", controllers.Hello)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	e.Logger.Fatal(e.Start(":" + port))
}
