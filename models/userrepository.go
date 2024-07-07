package models

import (
	"github.com/zykunov/timeTracker/storage"
)

func AddUser(u *User) error {
	return storage.DB.Create(u).Error
}

func DeleteUserById(u *User, id string) error {
	return storage.DB.Where("id = ?", id).Delete(u).Error
}

func GetUserById(u *User, id uint) error {
	return storage.DB.Where("id = ?", id).Error
}

func GetUserByPassport(u *User, passportSerie int, passportNumber int) error {
	return storage.DB.Where("passport_serie = ? AND passport_number = ?", passportSerie, passportNumber).First(u).Error
}

func UpdateUserById(u *User, id uint) error {
	return storage.DB.Updates(u).Error
}

func GetAllUsers(u *[]User, limit int, offset int) error {
	if err := storage.DB.Limit(limit).Offset(offset).Find(u).Error; err != nil {
		return err
	}
	return nil
}
