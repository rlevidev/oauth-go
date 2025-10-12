package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rlevidev/oauth-go/src/controllers"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.RouterGroup, db *gorm.DB) {
	v1 := r.Group("/api/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1.POST("/users/register", func(c *gin.Context) {
		controllers.CreateUser(c, db)
	})
	v1.POST("/users/login", func(c *gin.Context) {
		controllers.LoginUser(c, db)
	})
}
