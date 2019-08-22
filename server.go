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
