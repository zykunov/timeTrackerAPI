package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/zykunov/timeTracker/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	apiV1Group := router.Group("/api/v1")
	{
		apiV1Group.POST("useradd", handlers.AddUser)                   // Добавление пользователя
		apiV1Group.DELETE("userdelete/:passport", handlers.DeleteUser) // Удаление пользователя
		apiV1Group.PUT("article/:id", handlers.UpdateArticleById)      // Изменение пользователя
		apiV1Group.POST("start", handlers.StartTask)                   //Начать отсчет времени по задаче пользователя
		apiV1Group.POST("stop", handlers.StopTask)                     //Закончить отсчет времени по задаче пользователя
	}

	return router
}
