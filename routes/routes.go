package routes

import (
	"github.com/gin-gonic/gin"
	"restfultest/controllers"
)

func SetupRoutes(r *gin.Engine) {
	usersRoutes := r.Group("/users")
	{
		usersRoutes.POST("/", controllers.CreateUser)
		usersRoutes.GET("/", controllers.GetAllUsers)
	}
}
