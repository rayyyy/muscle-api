package main

import (
	"log"
	"muscle-api/controllers"
	"os"

	"muscle-api/funcs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(funcs.AuthHandler)

	e.POST("/sign-in", controllers.SignIn)
	e.GET("/user/:id", controllers.Get)
	e.POST("/user/:id", controllers.Update)
	e.GET("/hello", controllers.Hello)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	e.Logger.Fatal(e.Start(":" + port))
}
