package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zykunov/timeTracker/models"
	"github.com/zykunov/timeTracker/routers"
	"github.com/zykunov/timeTracker/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/zykunov/timeTracker/docs"
)

// @title           time tracker
// @version         1.0
// @description     REST API учета времени задач по пользователям

// @contact.name   Igor Zykunov

// @host      localhost:8080
// @BasePath  /api/v1

var err error

func init() {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	dbname := os.Getenv("DBNAME")
	password := os.Getenv("PASSWORD")
	sslmode := os.Getenv("SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=%s", host, user, dbname, password, sslmode)

	storage.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("error while accesing DB", err)
	}

	storage.DB.AutoMigrate(&models.User{}, &models.Task{})

	r := routers.SetupRouter()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
