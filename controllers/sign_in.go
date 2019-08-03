package controllers

import (
	"log"
	"net/http"

	"muscle-api/models"

	"github.com/labstack/echo/v4"
)

// User 認証用
type authUser struct {
	Email string `gorm:"column:email" json:"email"`
	UID   string `gorm:"column:uid" json:"uid"`
}

// SignIn サインイン用 ユーザーがいなかったら作成
func SignIn(c echo.Context) error {
	db, err := models.GetDB()
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	defer db.Close()

	authUser := new(authUser)
	if err := c.Bind(&authUser); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	user := new(models.User)
	db.Preload("UserDetail").Assign(authUser).FirstOrInit(&user, models.User{UID: authUser.UID})

	if db.NewRecord(user) == true {
		user.Nickname = "筋トレ初心者"
		user.Image = "exampleData"
	}
	if err := db.Save(&user).Error; err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, user)
}
