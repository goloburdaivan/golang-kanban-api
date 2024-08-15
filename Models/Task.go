package Models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	ID          int
	Title       string
	Description string
	Status      string
	AssigneeID  int
	CreatorID   int
	Assignee    User
	Creator     User
	Priority    int
	DueDate     time.Time
}
