package setup

import (
	"Golang/Http/Controllers"
	"Golang/Repository"
	"Golang/Services"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func BuildContainer(db *gorm.DB) *dig.Container {
	container := dig.New()

	// DB injection
	container.Provide(func() *gorm.DB {
		return db
	})

	// Repositories
	container.Provide(Repository.NewProjectRepository)
	container.Provide(Repository.NewUserRepository)
	container.Provide(Repository.NewTaskRepository)

	// Services
	container.Provide(Services.NewEmailConfirmationService)
	container.Provide(Services.NewUserService)
	container.Provide(Services.NewTaskService)

	// Controllers
	container.Provide(Controllers.NewUserController)
	container.Provide(Controllers.NewProjectController)
	container.Provide(Controllers.NewTaskController)

	return container
}
