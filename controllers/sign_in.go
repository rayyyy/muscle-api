package controllers

import (
	"log"
	"net/http"

	"muscle-api/funcs"
	"muscle-api/models"

	"github.com/k0kubun/pp"
	"github.com/labstack/echo/v4"
)

// AuthUser 認証用
type User struct {
	Email string `gorm:"column:email"`
	UID   string `gorm:"column:uid"`
}

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

	authUser := new(User)
	if err := c.Bind(&authUser); err != nil {
		return err
	}

	user := new(models.User)
	db.FirstOrInit(&user, authUser)

	if db.NewRecord(user) == true {
		user.Nickname = "筋トレ初心者"
		user.Image = "exampleData"
		if err := db.Save(&user).Error; err != nil {
			pp.Print(err)
		}
	}

	return c.JSON(http.StatusOK, user)
}
