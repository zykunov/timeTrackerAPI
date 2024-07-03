package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/zykunov/timeTracker/models"
	"github.com/zykunov/timeTracker/routers"
	"github.com/zykunov/timeTracker/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

	fmt.Println(dsn)
	storage.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("error while accesing DB", err)
	}

	storage.DB.AutoMigrate(&models.User{}, &models.Task{})

	r := routers.SetupRouter()
	r.Run()
}
