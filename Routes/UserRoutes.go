package Routes

import (
	"Golang/Http/Controllers"
	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine, userController Controllers.UserController) {
	routes := route.Group("/api/")
	{
		routes.POST("/login", userController.Login)
		routes.POST("/register", userController.Register)
	}
}
