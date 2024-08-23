package routes

import (
	"github.com/gin-gonic/gin"
	"restfultest/controllers"
)

func SetupRoutes(r *gin.Engine) {
	postRoutes := r.Group("/posts")
	{
		postRoutes.POST("/", controllers.CreatePost)
		postRoutes.GET("/", controllers.GetAllPosts)
	}
}
