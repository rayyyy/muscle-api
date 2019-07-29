package funcs

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// AuthHandler 認証しているか確認
func AuthHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		if isLogged := CheckLogin(idToken); isLogged == false {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		return next(c)
	}
}
