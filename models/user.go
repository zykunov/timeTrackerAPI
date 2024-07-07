package models

import (
	"time"

	_ "github.com/lib/pq"
)

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
