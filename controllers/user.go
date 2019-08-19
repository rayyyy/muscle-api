package controllers

import (
	"log"
	"muscle-api/models"
	"net/http"

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

	if err := db.Preload("UserDetail").Preload("Mentor").First(&user, id).Error; gorm.IsRecordNotFoundError(err) {
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

	updateUser := new(userParams)
	if err := c.Bind(&updateUser); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	user := new(models.User)
	db.First(&user, updateUser.ID)
	db.Model(&user).Updates(updateUser)

	db.Assign(updateUser.UserDetail).FirstOrInit(&user.UserDetail, updateUser.UserDetail.ID)
	db.Save(&user)

	return c.JSON(http.StatusOK, user)
}

type userParams struct {
	ID         int              `gorm:"column:id;primary_key" json:"id;primary_key"`
	Nickname   string           `gorm:"column:nickname" json:"nickname"`
	Birthday   null.Time        `gorm:"column:birthday" json:"birthday"`
	Sex        null.String      `gorm:"column:sex" json:"sex"`
	UserDetail userDetailParams `json:"user_detail"`
}
type userDetailParams struct {
	ID           int         `gorm:"column:id;primary_key" json:"id;primary_key"`
	UserID       int         `gorm:"column:user_id" json:"user_id"`
	ShortMessage null.String `gorm:"column:short_message" json:"short_message"`
	Height       null.Float  `gorm:"column:height" json:"height"`
	Weight       null.Float  `gorm:"column:weight" json:"weight"`
}
