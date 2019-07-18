package controllers

import (
	"log"
	"github.com/k0kubun/pp"
	"net/http"
	"muscle-api/models"

	"github.com/labstack/echo/v4"
)

// Update userの情報更新
func Update(c echo.Context) error {
	// name := c.Param("user_id") path param
	// https://echo.labstack.com/guide/request

// {
//   "id": 1,
//   "user_id": 1,
//   "short_message": "8ヶ月",
//   "height": 1,
//   "weight": 1
// }
	db, err := models.GetDB()
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Error!")
	}
	defer db.Close()

	userDetail := new(models.UserDetail)
	if err := c.Bind(&userDetail); err != nil {
		pp.Println(err)
		return err
	}
	pp.Println(userDetail)

	db.FirstOrInit(&userDetail, userDetail)
	db.Save(&userDetail)
	pp.Println(userDetail)

	return c.String(http.StatusOK, "Hello!")
}
