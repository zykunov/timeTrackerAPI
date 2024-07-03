package models

import (
	"github.com/zykunov/timeTracker/storage"
)

func AddUser(u *User) error {
	return storage.DB.Create(u).Error
}

func DeleteUserById(u *User, id int) error {
	return storage.DB.Where("passport = ?", id).Delete(u).Error
}
