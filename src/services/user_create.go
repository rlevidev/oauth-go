package services

import (
	resterr "github.com/rlevidev/oauth-go/src/config/rest_err"
	"github.com/rlevidev/oauth-go/src/models"
	"gorm.io/gorm"
)

func CreateUser(userDomain models.UserDomain, db *gorm.DB) (*models.UserDomain, *resterr.RestErr) {
	// Criptografar a senha
	userDomain.EncryptPassword()

	// Salvar usuário no banco de dados
	err := db.Create(&userDomain).Error
	if err != nil {
		return nil, resterr.NewInternalServerError("Erro ao salvar usuário no banco de dados")
	}

	return &userDomain, nil
}
