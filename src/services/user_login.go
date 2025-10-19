package services

import (
	resterr "github.com/rlevidev/oauth-go/src/config/jwt"
	resterrpkg "github.com/rlevidev/oauth-go/src/config/rest_err"
	"github.com/rlevidev/oauth-go/src/models"
	"gorm.io/gorm"
)

type LoginResponse struct {
	User  *models.UserDomain `json:"user"`
	Token string             `json:"token"`
}

func LoginUser(userDomain models.UserDomain, db *gorm.DB) (*LoginResponse, *resterrpkg.RestErr) {
	// Buscar usuário no banco pelo email
	var userFound models.UserDomain
	err := db.Where("email = ?", userDomain.Email).First(&userFound).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, resterrpkg.NewBadRequestError("Credenciais inválidas")
		}
		return nil, resterrpkg.NewInternalServerError("Erro ao buscar usuário")
	}

	// Criptografar a senha fornecida para comparar
	userDomain.EncryptPassword()

	// Verificar se a senha está correta
	if userFound.Password != userDomain.Password {
		return nil, resterrpkg.NewBadRequestError("Credenciais inválidas")
	}

	// Gerar token JWT
	token, err := resterr.GenerateToken(&userFound)
	if err != nil {
		return nil, resterrpkg.NewInternalServerError("Erro ao gerar token de autenticação")
	}

	loginResponse := &LoginResponse{
		User:  &userFound,
		Token: token,
	}

	return loginResponse, nil
}
