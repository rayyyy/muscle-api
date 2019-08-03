package controllers

import (
	"log"
	"muscle-api/models"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	null "gopkg.in/guregu/null.v3"
)

// Get userの情報取得
func Get(c echo.Context) error {
	db, err := models.GetDB()
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	defer db.Close()

	id := c.Param("id")
	user := new(models.User)
	db.Preload("UserDetail").First(&user, id)

	if err := db.Preload("UserDetail").First(&user, id).Error; gorm.IsRecordNotFoundError(err) {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, user)
}

// Update userの情報更新
func Update(c echo.Context) error {
	db, err := models.GetDB()
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	defer db.Close()

	updateUser := new(updateUser)
	if err := c.Bind(&updateUser); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	user := new(models.User)
	db.First(&user)
	db.Model(&user).Updates(updateUser)

	db.Assign(updateUser.UserDetail).FirstOrInit(&user.UserDetail)
	db.Save(&user)

	return c.JSON(http.StatusOK, user)
}

type updateUser struct {
	ID         int              `gorm:"column:id;primary_key" json:"id;primary_key"`
	Nickname   string           `gorm:"column:nickname" json:"nickname"`
	UserDetail updateUserDetail `json:"user_detail"`
}
type updateUserDetail struct {
	Height       null.Float  `gorm:"column:height" json:"height"`
	ID           int         `gorm:"column:id;primary_key" json:"id;primary_key"`
	ShortMessage null.String `gorm:"column:short_message" json:"short_message"`
	UserID       int         `gorm:"column:user_id" json:"user_id"`
	Weight       null.Float  `gorm:"column:weight" json:"weight"`
}
