package Routes

import (
	"Golang/Http/Controllers"
	"Golang/Http/Middleware"
	"github.com/gin-gonic/gin"
)

func Task(route *gin.Engine, taskController Controllers.TaskController) {
	routes := route.Group("/api/tasks", Middleware.RequiresAuth())
	{
		routes.GET("/", taskController.Index)
		routes.GET("/:id", taskController.Show)
		routes.POST("/", taskController.Create)
		routes.PUT("/:id", taskController.Update)
		routes.DELETE("/:id", taskController.Destroy)
		routes.POST("/:id/move", taskController.Move)
	}
}
