package controllers

import (
	"log"
	"muscle-api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UpdateMentorPlan メンタープランの情報更新
func UpdateMentorPlan(c echo.Context) error {
	db, err := models.GetDB()
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	defer db.Close()

	updateMentor := new(updateMentorParams)
	if err := c.Bind(&updateMentor); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	id := c.Param("id")
	user := new(models.User)
	db.First(&user, id)

	db.Assign(updateMentor).FirstOrInit(&user.Mentor, models.Mentor{UserID: user.ID})
	db.Save(&user)

	for _, plan := range updateMentor.MentorPlans {
		var mentorPlan models.MentorPlan
		plan.IsValid = 1 // frontから飛ばせばいらない気がする
		plan.MentorID = user.Mentor.ID
		db.Assign(plan).FirstOrInit(&mentorPlan, plan.ID)
		db.Save(&mentorPlan)
	}

	return c.JSON(http.StatusOK, user)
}

type updateMentorParams struct {
	ID            int                `json:"id"`
	Title         string             `json:"title"`
	AppealMessage string             `json:"appeal_message"`
	MentorPlans   []mentorPlanParams `json:"mentor_plans"`
}

type mentorPlanParams struct {
	ID       int    `json:"id"`
	MentorID int    `json:"mentor_id"`
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	Type     string `json:"type"`
	Price    int    `json:"price"`
	IsValid  int    `json:"is_valid"`
}