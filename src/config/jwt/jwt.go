package jwt

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rlevidev/oauth-go/src/models"
)

func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("FATAL: JWT_SECRET não foi definida no arquivo .env. A aplicação não pode iniciar sem uma chave secreta.")
	}
	return secret
}

func getJWTSecretBytes() []byte {
	return []byte(getJWTSecret())
}

// Definir dados do token
type Claims struct {
	UserID               string `json:"user_id"`
	Email                string `json:"email"`
	Name                 string `json:"name"`
	jwt.RegisteredClaims        // Expiração, emissão, etc.
}

// Gerar token
func GenerateToken(ud *models.UserDomain) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token válido por 24 horas

	claims := &Claims{
		UserID: ud.ID,
		Email:  ud.Email,
		Name:   ud.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "oauth-go",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Cria o token sem assinatura
	tokenString, err := token.SignedString(getJWTSecretBytes()) // Assina o token
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Verificar token
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return getJWTSecretBytes(), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

// Extrair token do cabeçalho
func ExtractTokenFromAuthHeader(authHeader string) string {
	const bearerPrefix = "Bearer "
	if len(authHeader) > len(bearerPrefix) && authHeader[:len(bearerPrefix)] == bearerPrefix {
		return authHeader[len(bearerPrefix):]
	}
	return ""
}
