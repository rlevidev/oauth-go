package services

import (
	resterr "github.com/rlevidev/oauth-go/src/config/rest_err"
	"github.com/rlevidev/oauth-go/src/models"
)

func CreateUser(userDomain models.UserDomain) (*models.UserDomain, *resterr.RestErr) {
	userDomain.EncryptPassword()

	return &userDomain, nil
}
