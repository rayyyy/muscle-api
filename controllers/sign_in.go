package controllers

import (
	"log"
	"net/http"

	"muscle-api/models"
	"muscle-api/funcs"

	"github.com/k0kubun/pp"
	"github.com/labstack/echo/v4"
)

// SignIn サインイン用 ユーザーがいなかったら作成
func SignIn(c echo.Context) error {
	if isLoggedIn := funcs.IsLoggedIn(c); isLoggedIn == false {
		return nil
	}

	db, err := models.GetDB()
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Error!")
	}
	defer db.Close()

	user := models.User{Nickname: "ray", UID: "abcdefghij", Image: "example", Email: "example@aaa.ccc"}
	pp.Print(user)
	if err := db.Create(&user).Error; err != nil {
		// エラーハンドリング...
		pp.Print(err)
	}

	return c.String(http.StatusOK, "SignIn!")
}
