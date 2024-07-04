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

func RespondJSON(w *gin.Context, status_code int, data interface{}) {
	log.Println("status code:", status_code)
	var message Message

	message.StatusCode = status_code
	message.Data = data
	w.JSON(200, message)
}
