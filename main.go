package main

import (
	"Golang/Http/Controllers"
	"Golang/Repository"
	"Golang/Routes"
	"Golang/Services"
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
		projectRepository        = Repository.NewProjectRepository(db)
		emailConfirmationService = Services.NewEmailConfirmationService()
		userRepository           = Repository.NewUserRepository(db)
		userService              = Services.NewUserService(userRepository, emailConfirmationService)
		userController           = Controllers.NewUserController(userRepository, emailConfirmationService, userService)
		projectsController       = Controllers.NewProjectController(projectRepository)
	)

	server := gin.Default()
	Routes.Project(server, projectsController)
	Routes.User(server, userController)
	server.Run()
}
