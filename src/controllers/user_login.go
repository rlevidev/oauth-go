package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rlevidev/oauth-go/src/config/validation"
	"github.com/rlevidev/oauth-go/src/controllers/request"
	"github.com/rlevidev/oauth-go/src/controllers/response"
	"github.com/rlevidev/oauth-go/src/models"
	"github.com/rlevidev/oauth-go/src/services"
	"gorm.io/gorm"
)

func LoginUser(ctx *gin.Context, db *gorm.DB) {
	var userLoginRequest request.UserLoginRequest

	if err := ctx.ShouldBindJSON(&userLoginRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		ctx.JSON(restErr.Status, restErr)
		return
	}

	userDomain := models.NewUserLoginDomain(
		userLoginRequest.Email,
		userLoginRequest.Password,
	)

	userDomainResult, err := services.LoginUser(*userDomain, db)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	userResponse := response.UserResponse{
		ID:    userDomainResult.ID,
		Email: userDomainResult.Email,
		Name:  userDomainResult.Name,
	}

	ctx.JSON(200, userResponse)
}
