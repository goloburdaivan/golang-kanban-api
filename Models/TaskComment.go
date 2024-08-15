package Models

import "gorm.io/gorm"

type TaskComment struct {
	gorm.Model
	ID      int
	TaskID  int
	Task    Task
	UserID  int
	User    User
	Content string
}
