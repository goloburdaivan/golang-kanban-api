package Services

import (
	"Golang/Models"
	"gorm.io/gorm"
)

func MigrateSchemas(db *gorm.DB) {
	db.AutoMigrate(&Models.User{})
	db.AutoMigrate(&Models.Project{})
	db.AutoMigrate(&Models.Task{})
	db.AutoMigrate(&Models.Column{})
	db.AutoMigrate(&Models.TaskComment{})
}
