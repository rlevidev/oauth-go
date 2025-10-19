package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rlevidev/oauth-go/src/controllers"
	"github.com/rlevidev/oauth-go/src/middleware"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.RouterGroup, db *gorm.DB) {
	v1 := r.Group("/api/v1")

	// Rota pública para verificar se o servidor está funcionando
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Rotas públicas de autenticação
	v1.POST("/users/register", func(c *gin.Context) {
		controllers.CreateUser(c, db)
	})
	v1.POST("/users/login", func(c *gin.Context) {
		controllers.LoginUser(c, db)
	})

	// Grupo de rotas protegidas que requerem autenticação
	protected := v1.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// Exemplo de rota protegida - obter perfil do usuário autenticado
		protected.GET("/users/profile", func(c *gin.Context) {
			userID, exists := c.Get("user_id")
			if !exists {
				c.JSON(401, gin.H{"error": "Usuário não autenticado"})
				return
			}

			userEmail, _ := c.Get("user_email")
			userName, _ := c.Get("user_name")

			c.JSON(200, gin.H{
				"user_id":  userID,
				"email":    userEmail,
				"name":     userName,
			})
		})

		// Adicione mais rotas protegidas aqui conforme necessário
		// protected.GET("/users", controllers.GetUsers)
		// protected.PUT("/users/:id", controllers.UpdateUser)
		// protected.DELETE("/users/:id", controllers.DeleteUser)
	}
}
