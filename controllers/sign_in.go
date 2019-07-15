package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// SignIn サインイン用
func SignIn(c echo.Context) error {
	// ここでユーザーログインする
	// ユーザーがいなかったら作成
	return c.String(http.StatusOK, "SignIn!")
}
