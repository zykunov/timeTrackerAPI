package helpers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zykunov/timeTracker/models"
)

type Message struct {
	StatusCode int         `json:"status_code"`
	Meta       interface{} `json:"meta"`
	Data       interface{} `json:"data"`
}

type GetWorkFinal struct {
	ID     int     `json:"userId"`
	TaskId int     `json:"taskId"`
	Hours  float64 `json:"time"`
}

func RespondJSON(w *gin.Context, status_code int, data interface{}) {
	log.Println("status code:", status_code)
	var message Message

	message.StatusCode = status_code
	message.Data = data
	w.JSON(200, message)
}

func GetFinalWork(t *[]models.Task) GetWorkFinal {

	return GetWorkFinal{}
}
