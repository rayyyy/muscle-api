package controllers

import (
	"log"
	"net/http"
	"muscle-api/models"

	"github.com/labstack/echo/v4"
)

// Update userの情報更新
func Update(c echo.Context) error {
	db, err := models.GetDB()
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Error!")
	}
	defer db.Close()

	userDetail := new(models.UserDetail)
	if err := c.Bind(&userDetail); err != nil {
		log.Println(err)
		return err
	}

	db.FirstOrInit(&userDetail, userDetail)
	db.Save(&userDetail)

	return c.JSON(http.StatusOK, userDetail)
}
