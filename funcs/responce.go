package funcs

import (
	"net/http"
	"strings"
	"github.com/labstack/echo/v4"
)

// IsLoggedIn ユーザーがログインしているかを返す
func IsLoggedIn(c echo.Context) bool {
	authHeader := c.Request().Header.Get("Authorization")
	idToken := strings.Replace(authHeader, "Bearer ", "", 1)
	if isLogged := CheckLogin(idToken); isLogged == false {
		c.String(http.StatusUnauthorized, "認証失敗")
		return false
	}
	return true
}
