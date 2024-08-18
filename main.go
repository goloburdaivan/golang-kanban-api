package main

import (
	"Golang/Http/Controllers"
	"Golang/Repository"
	"Golang/Routes"
	"Golang/setup"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	setup.LoadEnvironment()
	db = setup.InitDatabase()
	setup.MigrateDatabase(db)
}

func main() {
	var (
		projectRepository  = Repository.NewProjectRepository(db)
		userRepository     = Repository.NewUserRepository(db)
		userController     = Controllers.NewUserController(userRepository)
		projectsController = Controllers.NewProjectController(projectRepository)
	)

	server := gin.Default()
	Routes.Project(server, projectsController)
	Routes.User(server, userController)
	server.Run()
}
