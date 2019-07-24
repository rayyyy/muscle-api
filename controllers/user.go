package controllers

import (
	"log"
	"muscle-api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Get userの情報取得
func Get(c echo.Context) error {
	db, err := models.GetDB()
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Error!")
	}
	defer db.Close()

	id := c.Param("id")
	var user models.User
	db.Preload("UserDetail").First(&user, id)

	return c.JSON(http.StatusOK, user)
}

// Update userの情報更新
func Update(c echo.Context) error {
	db, err := models.GetDB()
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Error!")
	}
	defer db.Close()

	user := new(models.User)
	if err := c.Bind(&user); err != nil {
		log.Println(err)
		return err
	}

	db.FirstOrInit(&user, user)
	db.Save(&user)

	return c.JSON(http.StatusOK, user)
}
