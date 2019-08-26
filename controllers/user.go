package controllers

import (
	"log"
	"muscle-api/funcs"
	"muscle-api/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	null "gopkg.in/guregu/null.v3"
)

// GetUser userの情報取得
func GetUser(c echo.Context) error {
	db, err := models.GetDB()
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	defer db.Close()

	id := c.Param("id")
	user := new(models.User)
	db.First(&user, id).Related(&user.UserDetail)

	if err := db.Set("gorm:auto_preload", true).First(&user, id).Error; gorm.IsRecordNotFoundError(err) {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, user)
}

// UpdateUser userの情報更新
func UpdateUser(c echo.Context) error {
	db, err := models.GetDB()
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	defer db.Close()

	var params struct {
		User     userParams `json:"user"`
		UserIcon string     `json:"user_icon"`
	}
	if err := c.Bind(&params); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	updateUser := params.User

	if params.UserIcon != "" {
		success := funcs.UploadImage(strconv.Itoa(updateUser.ID)+".png", strings.Split(params.UserIcon, "base64,")[1])
		if success == true {
			updateUser.Image = "https://storage.googleapis.com/muscle-2710.appspot.com/user-icon/" + strconv.Itoa(updateUser.ID) + ".png"
		}
	}

	user := new(models.User)
	db.First(&user, updateUser.ID)
	db.Model(&user).Updates(updateUser)

	db.Assign(updateUser.UserDetail).FirstOrInit(&user.UserDetail, updateUser.UserDetail.ID)
	db.Save(&user)

	return c.JSON(http.StatusOK, user)
}

type userParams struct {
	ID         int              `json:"id"`
	Image      string           `json:"image"`
	Nickname   string           `json:"nickname"`
	Birthday   null.Time        `json:"birthday"`
	Sex        null.String      `json:"sex"`
	UserDetail userDetailParams `json:"user_detail"`
}
type userDetailParams struct {
	ID           int         `json:"id"`
	UserID       int         `json:"user_id"`
	ShortMessage null.String `json:"short_message"`
	Height       null.Float  `json:"height"`
	Weight       null.Float  `json:"weight"`
}
