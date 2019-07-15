package controllers

import (
	"log"
	"net/http"

	"muscle-api/models"

	"github.com/labstack/echo/v4"
)

// SignIn サインイン用
func SignIn(c echo.Context) error {
	db, err := models.GetDB()
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "Error!")
	}
	defer db.Close()

	// ここでユーザーログインする
	// ユーザーがいなかったら作成
	return c.String(http.StatusOK, "SignIn!")
}
