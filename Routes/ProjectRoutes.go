package Routes

import (
	"Golang/Http/Controllers"
	"Golang/Http/Middleware"
	"github.com/gin-gonic/gin"
)

func Project(route *gin.Engine, projectController Controllers.ProjectController) {
	routes := route.Group("/api/project", Middleware.RequiresAuth())
	{
		routes.GET("/", projectController.Index)
		routes.GET("/:id", projectController.Show)
		routes.POST("/", Middleware.RequiresRole("admin"), projectController.Create)
		routes.PUT("/:id", Middleware.RequiresRole("admin"), projectController.Update)
		routes.DELETE("/:id", Middleware.RequiresRole("admin"), projectController.Destroy)
	}
}
