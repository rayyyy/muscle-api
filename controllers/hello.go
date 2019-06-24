package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Hello お試し用の関数
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello!")
}
