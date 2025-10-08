package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rlevidev/oauth-go/src/controllers"
)

func InitRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/api/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1.POST("/users/register", controllers.CreateUser)
}
