package main

import (
	"log"
	"muscle-api/controllers"
	"muscle-api/handler"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200", "https://muscle.netlify.com"},
	}))
	e.Use(handler.AuthHandler)

	e.POST("/sign-in", controllers.SignIn)
	e.GET("/user/:id", controllers.Get)
	e.POST("/user/:id", controllers.Update)
	e.GET("/hello", controllers.Hello)

	e.HTTPErrorHandler = handler.JSONErrorHandler

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	e.Logger.Fatal(e.Start(":" + port))
}
