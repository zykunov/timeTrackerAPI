package helpers

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Message struct {
	StatusCode int         `json:"status_code"`
	Meta       interface{} `json:"meta"`
	Data       interface{} `json:"data"`
}

type GetWork struct {
	ID        int    `json:"userId"`
	DateStart string `json:"dateStart"`
	DateEnd   string `json:"dateEnd"`
}

type TaskStartStop struct {
	ID     int `json:"userId"`
	TaskId int `json:"taskId"`
}

// @Description limit
// @Description offset
type Paging struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// @Description passportNumber
// @Description passportSerie
type Passport struct {
	PassportNumber string `json:"passportnumber"`
	PassportSerie  string `json:"passportserie"`
}

type UserAddStruct struct {
	PassportNumber string `json:"passportNumber"`
}

type UserUpdate struct {
	ID             uint   `json:"ID"` //@Description required
	PassportSerie  int    `json:"passportSerie"`
	PassportNumber int    `json:"passportNumber"`
	Surname        string `json:"surname"`
	Name           string `json:"name"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
}

type GetWorkFinal struct {
	ID       uint    `json:"Id"`
	UserId   int     `json:"userId"`
	TaskTime float64 `json:"taskTime"`
}

func RespondJSON(w *gin.Context, status_code int, data interface{}) {
	log.Println("status code:", status_code)
	var message Message

	message.StatusCode = status_code
	message.Data = data
	w.JSON(200, message)
}
