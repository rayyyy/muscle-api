package main

import (
	"muscle-api/funcs"
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

	var allowOrigins []string
	if funcs.IsProduction() {
		allowOrigins = append(allowOrigins, "https://muscle.netlify.com")
	} else {
		allowOrigins = append(allowOrigins, "*")
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: allowOrigins,
	}))

	e.GET("/user/:id", controllers.GetUser)

	// 認証を必要とするもの
	g := e.Group("/auth", handler.AuthHandler)
	g.POST("/sign-in", controllers.SignIn)
	g.POST("/user/:id", controllers.UpdateUser)
	g.POST("/user/:id/mentor-plan", controllers.UpdateMentorPlan)

	e.HTTPErrorHandler = handler.JSONErrorHandler

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	e.Logger.Fatal(e.Start(":" + port))
}
