package services

import (
	resterr "github.com/rlevidev/oauth-go/src/config/rest_err"
	"github.com/rlevidev/oauth-go/src/models"
	"gorm.io/gorm"
)

func CreateUser(userDomain models.UserDomain, db *gorm.DB) (*models.UserDomain, *resterr.RestErr) {
	// Verificar se já existe um usuário com este email
	var existingUser models.UserDomain
	err := db.Where("email = ?", userDomain.Email).First(&existingUser).Error

	if err == nil {
		// Usuário encontrado - email já existe
		return nil, resterr.NewBadRequestError("Este email já está cadastrado no sistema")
	}

	if err != gorm.ErrRecordNotFound {
		// Erro diferente de "não encontrado" - problema no banco
		return nil, resterr.NewInternalServerError("Erro ao verificar email existente")
	}

	// Email não existe, pode prosseguir com a criação
	
	// Criptografar a senha
	userDomain.EncryptPassword()

	// Salvar usuário no banco de dados
	err = db.Create(&userDomain).Error
	if err != nil {
		return nil, resterr.NewInternalServerError("Erro ao salvar usuário no banco de dados")
	}

	return &userDomain, nil
}
