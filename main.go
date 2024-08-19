package main

import (
	"Golang/Http/Controllers"
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

	diContainer := setup.BuildContainer(db)

	server := gin.Default()
	diContainer.Invoke(func(
		userController Controllers.UserController,
		projectsController Controllers.ProjectController,
		taskController Controllers.TaskController,
	) {
		Routes.Project(server, projectsController)
		Routes.User(server, userController)
		Routes.Task(server, taskController)
	})

	server.Run()
}
