package Models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	ID          int
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      int    `json:"user_id"`
	User        *User  `json:"owner,omitempty"`
}
