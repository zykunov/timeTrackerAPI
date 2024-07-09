package handlers

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zykunov/timeTracker/helpers"
	"github.com/zykunov/timeTracker/models"
)

// @Summary AddUser
// @Tags user
// @Description Добавление пользователя. Пасппортные данные - строка, разделитель пробел.
// @Accept json
// @Produce json
// @Param input body helpers.UserAddStruct true "user info"
// @Success 201 {object} models.User
// @Failure 404 {object} models.User
// @Router /useradd [post]
func AddUser(c *gin.Context) {

	var passport struct {
		Raw string `json:"passportNumber"`
	}

	if c.BindJSON(&passport) != nil {
		c.String(400, "parameter error")
		return
	}

	explode := strings.Split(passport.Raw, " ")
	passportNumber, err := strconv.Atoi(explode[0])
	if err != nil {
		c.String(400, "wrong passport")
		return
	}
	passportSerie, err := strconv.Atoi(explode[1])
	if err != nil {
		c.String(400, "wrong passport")
		return
	}

	var user models.User

	user.PassportNumber = passportNumber
	user.PassportSerie = passportSerie

	err = models.AddUser(&user)
	if err != nil {
		helpers.RespondJSON(c, 404, user)
		return
	}
	helpers.RespondJSON(c, 200, user)
	log.Println("user added")
}

// StartTask             godoc
// @Summary      Start time count
// @Description  Начало отсчета времени по пользователю и задаче.
// @Tags         tasks
// @Produce      json
// @Param        task  body      helpers.TaskStartStop  true  "Task JSON"
// @Success      200   {object}  models.Task
// @Router       /start [post]
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
	helpers.RespondJSON(c, 200, task)
	log.Println("task started")

}

// StopTask             godoc
// @Summary      Stop time count
// @Description  На вход: Id пользователя и задачи, для которых закончить отсчет времени.
// @Tags         tasks
// @Produce      json
// @Param        task  body      helpers.TaskStartStop   true  "Task JSON"
// @Success      200   {object}  models.Task
// @Router       /stop [post]
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
	task.TaskTime = math.Floor(((float64(task.TaskEnd)-float64(task.TaskStart))/3600)*100) / 100

	err = models.UpdateTaskById(&task, task.ID)
	if err != nil {
		helpers.RespondJSON(c, 404, task)
	}
	helpers.RespondJSON(c, 200, task)
	log.Println("task ended")
}

// GetWork             godoc
// @Summary      Get work hours
// @Description  Получение трудозатрат по пользователю за период задача-сумма часов и минут с сортировкой от большей затраты к меньшей. Пример даты 2024-07-06
// @Tags         tasks
// @Produce      json
// @Param        task  body     helpers.GetWork true  "Task JSON"
// @Success      200   {object}  models.Task
// @Router       /getwork [Post]
func GetWork(c *gin.Context) {
	var tasks []models.Task
	var getWork helpers.GetWork

	if c.BindJSON(&getWork) != nil {
		c.String(400, "parameter error")
		return
	}

	dateStart, e := time.Parse(time.RFC3339, getWork.DateStart+"T22:08:41+00:00")
	if e != nil {
		panic("Parse error")
	}
	timestampStart := dateStart.Unix()

	dateEnd, e := time.Parse(time.RFC3339, getWork.DateEnd+"T22:08:41+00:00")
	if e != nil {
		panic("Parse error")
	}
	timestampEnd := dateEnd.Unix()

	err := models.GetWorkById(&tasks, getWork.ID, timestampStart, timestampEnd)
	if err != nil {
		helpers.RespondJSON(c, 404, tasks)
		return
	}

	helpers.RespondJSON(c, 200, tasks)
	log.Println("get work hours command")
}

// DeleteUser             godoc
// @Summary      User delete
// @Description  Удаление пользователя
// @Tags         user
// @Produce      json
// @Param id  path int true "User ID"
// @Success      200   {object}  models.User
// @Router /userdelete/{id}  [delete]
func DeleteUser(c *gin.Context) {
	var user models.User

	id := c.Params.ByName("id")

	err := models.DeleteUserById(&user, id)
	if err != nil {
		helpers.RespondJSON(c, 404, user)
		return
	}
	helpers.RespondJSON(c, 200, user)
	log.Println("user delete command")

}

// UserUpdate             godoc
// @Summary      User update
// @Description  Изменение пользователя
// @Tags         user
// @Produce      json
// @Param        user  body      helpers.UserUpdate  true  "User JSON"
// @Success      200   {object}  models.User
// @Router       /userupdate [patch]
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
	helpers.RespondJSON(c, 200, user)
	log.Println("user update command")
}

// GetUser             godoc
// @Summary      Get info about one user
// @Description  Получение данных пользователя по серии иномеру паспорта
// @Tags         user
// @Produce      json
// @Param        user  query   helpers.Passport  true  "User JSON"
// @Success      200  {object}  models.User
// @Router       /info [get]
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
	log.Println("get user command")
}

// GetUsers             godoc
// @Summary      Get info about users
// @Description  Получение всех пользователей, доступны параметры limit и offset ()
// @Tags         user
// @Produce      json
// @Param        user  query   helpers.Paging  true  "User JSON"
// @Success      200  {array}  models.User
// @Router       /getusers [get]
func GetUsers(c *gin.Context) {
	var user []models.User

	limit := c.Query("limit")
	offset := c.Query("offset")
	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	log.Println("limit:", limitInt)
	log.Println("offset:", offsetInt)

	err := models.GetAllUsers(&user, limitInt, offsetInt)

	if err != nil {
		helpers.RespondJSON(c, 400, user)
	}
	helpers.RespondJSON(c, 200, user)
	log.Println("get users command")
}
