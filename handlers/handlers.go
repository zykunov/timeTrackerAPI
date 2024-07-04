package handlers

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zykunov/timeTracker/helpers"
	"github.com/zykunov/timeTracker/models"
)

// @Summary AddUser
// @Tags User
// @Description user add function
// @Accept json
// @Produce json
// @Param input body models.User true "user info"
// @Success 201 {object} models.User
// @Failure 404 {object} models.User
// @Router /useradd [post]
func AddUser(c *gin.Context) {

	var user models.User
	if c.BindJSON(&user) != nil {
		c.String(400, "parameter error")
		return
	}

	err := models.AddUser(&user)
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

func DeleteUser(c *gin.Context) {
	var user models.User

	err := models.DeleteUserById(&user, user.ID)
	if err != nil {
		helpers.RespondJSON(c, 404, user)
		return
	}
	helpers.RespondJSON(c, 202, user)
}

func UpdateUserById(c *gin.Context) {
	var user models.User

	if c.BindJSON(&user) != nil {
		c.String(400, "parameter error")
		return
	}

	err := models.GetUserById(&user, user.ID)
	if err != nil {
		helpers.RespondJSON(c, 404, user)
		return
	}

	err = models.UpdateUserById(&user, user.ID)
	if err != nil {
		helpers.RespondJSON(c, 404, user)
	}
	helpers.RespondJSON(c, 202, user)
}

func GetUser(c *gin.Context) {
	var user models.User

	passportNumber := c.Query("passportnumber")

	passportSerie, ok := c.GetQuery("passportserie")
	if !ok {
		fmt.Println("Parameter does not exist")
		return
	}

	serie, _ := strconv.Atoi(passportSerie)
	number, _ := strconv.Atoi(passportNumber)

	log.Println(serie)
	log.Println(number)

	err := models.GetUserByPassport(&user, serie, number)

	if err != nil {
		helpers.RespondJSON(c, 400, user)
	}
	helpers.RespondJSON(c, 200, user)

}
