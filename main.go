package main

import (
	"Golang/Http/Controllers"
	"Golang/Repository"
	"Golang/Routes"
	"Golang/Services"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Services.MigrateSchemas(db)

	var (
		projectRepository  Repository.ProjectRepository  = Repository.NewProjectRepository(db)
		projectsController Controllers.ProjectController = Controllers.NewProjectController(projectRepository)
	)

	server := gin.Default()
	Routes.Project(server, projectsController)
	server.Run()
}
