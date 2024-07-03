package models

import (
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID        uint  `json:"taskID"`
	TaskStart int64 `json:"taskStarted"`
	TaskEnd   int64 `json:"taskEnded"`
	UserId    int   `json:"passportNumber"`
}

func (t *Task) TableName() string {
	return "task"
}
