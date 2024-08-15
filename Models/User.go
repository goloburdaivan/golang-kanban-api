package Models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           int
	Username     string
	Email        string
	PasswordHash string
}
