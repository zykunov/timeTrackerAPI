package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zykunov/timeTracker/handlers"

	_ "github.com/zykunov/timeTracker/docs"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	apiV1Group := router.Group("/api/v1")
	{
		apiV1Group.POST("useradd", handlers.AddUser)                   // Добавление пользователя
		apiV1Group.DELETE("userdelete/:passport", handlers.DeleteUser) // Удаление пользователя
		apiV1Group.PUT("userupdate", handlers.UpdateUserById)          // Изменение пользователя
		apiV1Group.POST("start", handlers.StartTask)                   //Начать отсчет времени по задаче пользователя
		apiV1Group.POST("stop", handlers.StopTask)                     //Закончить отсчет времени по задаче пользователя
		apiV1Group.GET("info", handlers.GetUser)                       //Получение пользователя
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
