package services

import (
	resterr "github.com/rlevidev/oauth-go/src/config/rest_err"
	"github.com/rlevidev/oauth-go/src/models"
	"gorm.io/gorm"
)

func LoginUser(userDomain models.UserDomain, db *gorm.DB) (*models.UserDomain, *resterr.RestErr) {
	// Buscar usuário no banco pelo email
	var userFound models.UserDomain
	err := db.Where("email = ?", userDomain.Email).First(&userFound).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, resterr.NewBadRequestError("Credenciais inválidas")
		}
		return nil, resterr.NewInternalServerError("Erro ao buscar usuário")
	}

	// Criptografar a senha fornecida para comparar
	userDomain.EncryptPassword()

	// Verificar se a senha está correta
	if userFound.Password != userDomain.Password {
		return nil, resterr.NewBadRequestError("Credenciais inválidas")
	}

	return &userFound, nil
}
