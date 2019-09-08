package controllers

import (
	"log"
	"muscle-api/funcs"
	"muscle-api/models"
	"net/http"
	"strconv"

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

	var params struct {
		Mentor updateMentorParams `json:"mentor"`
		Image1 string             `json:"image1"`
		Image2 string             `json:"image2"`
		Image3 string             `json:"image3"`
		Image4 string             `json:"image4"`
	}

	if err := c.Bind(&params); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	id := c.Param("id")
	user := new(models.User)
	db.First(&user, id)

	// DB失敗したらデータ消したい
	if params.Image1 != "" {
		success := funcs.UploadPlanImage(strconv.Itoa(user.ID)+"1.png", params.Image1)
		if success == true {
			params.Mentor.Image1 = "https://storage.googleapis.com/muscle-2710.appspot.com/mentor-image/" + strconv.Itoa(user.ID) + "1.png"
		}
	}

	if params.Image2 != "" {
		success := funcs.UploadPlanImage(strconv.Itoa(user.ID)+"2.png", params.Image2)
		if success == true {
			params.Mentor.Image2 = "https://storage.googleapis.com/muscle-2710.appspot.com/mentor-image/" + strconv.Itoa(user.ID) + "2.png"
		}
	}

	if params.Image3 != "" {
		success := funcs.UploadPlanImage(strconv.Itoa(user.ID)+"3.png", params.Image3)
		if success == true {
			params.Mentor.Image3 = "https://storage.googleapis.com/muscle-2710.appspot.com/mentor-image/" + strconv.Itoa(user.ID) + "3.png"
		}
	}

	if params.Image4 != "" {
		success := funcs.UploadPlanImage(strconv.Itoa(user.ID)+"4.png", params.Image4)
		if success == true {
			params.Mentor.Image4 = "https://storage.googleapis.com/muscle-2710.appspot.com/mentor-image/" + strconv.Itoa(user.ID) + "4.png"
		}
	}

	db.Assign(params.Mentor).FirstOrInit(&user.Mentor, models.Mentor{UserID: user.ID})
	db.Save(&user)

	for _, plan := range params.Mentor.MentorPlans {
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
	Image1        string             `json:"image1"`
	Image2        string             `json:"image2"`
	Image3        string             `json:"image3"`
	Image4        string             `json:"image4"`
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
