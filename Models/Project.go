package Models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	ID          int
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
