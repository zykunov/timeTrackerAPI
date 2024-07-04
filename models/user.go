package models

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/zykunov/timeTracker/storage"
)

func GetAllUsers(u *User) error {
	if err := storage.DB.Find(u).Error; err != nil {
		return err
	}
	return nil
}

type User struct {
	ID             uint `json:"ID"; gorm:"primarykey"`
	PassportSerie  int  `json:"passportSerie"`
	PassportNumber int  `json:"passportNumber"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Surname        string `json:"surname"`
	Name           string `json:"name"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
}

func (u *User) TableName() string {
	return "people"
}
