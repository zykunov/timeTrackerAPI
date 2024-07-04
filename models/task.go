package models

import (
	"time"

	_ "github.com/lib/pq"
)

type Task struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        uint  `json:"taskID"; gorm:"primarykey"`
	TaskStart int64 `json:"taskStarted"`
	TaskEnd   int64 `json:"taskEnded"`
	UserId    int   `json:"passportNumber"`
}

func (t *Task) TableName() string {
	return "task"
}
