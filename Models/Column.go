package Models

import "gorm.io/gorm"

type Column struct {
	gorm.Model
	ID        int
	Name      string
	Position  int
	ProjectID int
	Project   Project
}
