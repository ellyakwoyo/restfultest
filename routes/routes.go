package routes

import (
	"github.com/gin-gonic/gin"
	"restfultest/controllers"
)

func SetupRoutes(r *gin.Engine) {
	usersRoutes := r.Group("/users")
	{
		usersRoutes.GET("/", controllers.GetAllUsers)
		usersRoutes.GET("/:id", controllers.GetUser)
		usersRoutes.POST("/", controllers.CreateUser)
		usersRoutes.PUT("/:id", controllers.UpdateUser)
		usersRoutes.DELETE("/:id", controllers.DeleteUser)

	}
}
