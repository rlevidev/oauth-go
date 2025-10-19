package middleware

import (
	"github.com/gin-gonic/gin"
	jwtconfig "github.com/rlevidev/oauth-go/src/config/jwt"
	resterr "github.com/rlevidev/oauth-go/src/config/rest_err"
)

// Principal função de middleware para proteger rotas. Requer autenticação ( Bearer token )
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Pega o HEADER de autorização
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			restErr := resterr.NewUnauthorizedError("HEADER de autorização ausente")
			ctx.JSON(restErr.Status, restErr)
			ctx.Abort() // Interrompe a requisição
			return
		}

		// Extrai o token do HEADER
		tokenString := jwtconfig.ExtractTokenFromAuthHeader(authHeader)
		if tokenString == "" {
			restErr := resterr.NewUnauthorizedError("Token Bearer malformado")
			ctx.JSON(restErr.Status, restErr)
			ctx.Abort()
			return
		}

		// Valida o token
		claims, err := jwtconfig.ValidateToken(tokenString)
		if err != nil {
			restErr := resterr.NewUnauthorizedError("Token inválido ou expirado")
			ctx.JSON(restErr.Status, restErr)
			ctx.Abort()
			return
		}

		// Adicionar informações do usuário ao contexto
		ctx.Set("user_id", claims.UserID)
		ctx.Set("user_email", claims.Email)
		ctx.Set("user_name", claims.Name)

		// Continuar com a requisição
		ctx.Next()
	}
}

// Middleware opcional para verificar se o usuário está autenticado
func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader != "" {
			tokenString := jwtconfig.ExtractTokenFromAuthHeader(authHeader)
			if tokenString != "" {
				if claims, err := jwtconfig.ValidateToken(tokenString); err == nil {
					ctx.Set("user_id", claims.UserID)
					ctx.Set("user_email", claims.Email)
					ctx.Set("user_name", claims.Name)
				}
			}
		}
		ctx.Next()
	}
}
