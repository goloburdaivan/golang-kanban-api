package Routes

import (
	"Golang/Http/Controllers"
	"github.com/gin-gonic/gin"
)

func Project(route *gin.Engine, projectController Controllers.ProjectController) {
	routes := route.Group("/api/project")
	{
		routes.GET("/", projectController.Index)
		routes.GET("/:id", projectController.Show)
		routes.POST("/", projectController.Create)
		routes.PUT("/:id", projectController.Update)
		routes.DELETE("/:id", projectController.Destroy)
	}
}
