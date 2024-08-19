package Models

import "gorm.io/gorm"

type Column struct {
	gorm.Model
	Name      string `json:"name"`
	Position  int    `json:"position"`
	ProjectID int    `json:"project_id"`
	Project   Project
}
