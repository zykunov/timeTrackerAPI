package handlers

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zykunov/timeTracker/helpers"
	"github.com/zykunov/timeTracker/models"
)

func AddUser(c *gin.Context) {

	var PassportStruct struct {
		Passport string `json:"passportNumber"`
	}

	if c.BindJSON(&PassportStruct) != nil {
		c.String(400, "parameter error")
		return
	}

	//Перевод passportNumber из строки в int
	passportToTrim := strings.ReplaceAll(PassportStruct.Passport, " ", "")
	passportToInt, err := strconv.Atoi(passportToTrim)
	if err != nil {
		log.Println("error with ATOI")
	}

	var user models.User
	user.Passport = passportToInt

	err = models.AddUser(&user)
	if err != nil {
		helpers.RespondJSON(c, 404, user)
		return
	}
	helpers.RespondJSON(c, 201, user)
}

func StartTask(c *gin.Context) {

	var task models.Task

	if c.BindJSON(&task) != nil {
		c.String(400, "parameter error")
		return
	}

	task.TaskStart = time.Now().Unix()

	err := models.StartTaskFunc(&task)
	if err != nil {
		helpers.RespondJSON(c, 404, task)
		return
	}
	helpers.RespondJSON(c, 201, task)
}

func StopTask(c *gin.Context) {
	var task models.Task

	if c.BindJSON(&task) != nil {
		c.String(400, "parameter error")
		return
	}

	err := models.GetTaskById(&task, task.UserId, task.ID)
	if err != nil {
		helpers.RespondJSON(c, 404, task)
	}

	task.TaskEnd = time.Now().Unix()

	err = models.UpdateTaskById(&task, task.ID)
	if err != nil {
		helpers.RespondJSON(c, 404, task)
	}
	helpers.RespondJSON(c, 202, task)

}

func UpdateArticleById(c *gin.Context) {
	var article models.Article
	id := c.Params.ByName("id")
	err := models.GetArticleById(&article, id)
	if err != nil {
		helpers.RespondJSON(c, 404, article)
	}
	c.BindJSON(&article)
	err = models.UpdateArticleById(&article, id)
	if err != nil {
		helpers.RespondJSON(c, 404, article)
	}
	helpers.RespondJSON(c, 202, article)

}

func DeleteUser(c *gin.Context) {
	var user models.User
	passport := c.Params.ByName("passport")

	passportToInt, err := strconv.Atoi(passport)
	if err != nil {
		log.Println("error with ATOI")
	}

	err = models.DeleteUserById(&user, passportToInt)
	if err != nil {
		helpers.RespondJSON(c, 404, user)
		return
	}
	helpers.RespondJSON(c, 202, user)
}
