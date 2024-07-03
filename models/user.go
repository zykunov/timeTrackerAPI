package models

import (
	_ "github.com/lib/pq"
	"github.com/zykunov/timeTracker/storage"
	"gorm.io/gorm"
)

// дописать потом
func GetAllUsers(a *[]Article) error {
	if err := storage.DB.Find(a).Error; err != nil {
		return err
	}
	return nil

}

type User struct {
	gorm.Model
	Passport   int    `json:"passportNumber"`
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
	Address    string `json:"address"`
}

func (u *User) TableName() string {
	return "people"
}
