package Models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Status      string    `json:"status" binding:"required"`
	AssigneeID  int       `json:"assignee_id" binding:"required"`
	CreatorID   int       `json:"creator_id" binding:"required"`
	Priority    int       `json:"priority" binding:"required"`
	DueDate     time.Time `json:"due_date" binding:"required"`
	Assignee    User
	Creator     User
}
