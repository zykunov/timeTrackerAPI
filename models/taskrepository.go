package models

import (
	"github.com/zykunov/timeTracker/storage"
)

func StartTaskFunc(t *Task) error {
	return storage.DB.Create(t).Error
}

func GetTaskById(t *Task, userID int, taskId uint) error {
	return storage.DB.Where("id = ? AND user_id = ?", taskId, userID).First(t).Error
}

func UpdateTaskById(t *Task, id uint) error {
	return storage.DB.Updates(t).Error
}

func GetWorkById(t *[]Task, id int, dateStart int64, dateEnd int64) error {
	return storage.DB.Where("user_id = ? AND task_start > ? AND task_end < ?", id, dateStart, dateEnd).Order("task_time DESC").Find(t).Error
}
