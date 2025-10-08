package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rlevidev/oauth-go/src/config/validation"
	"github.com/rlevidev/oauth-go/src/controllers/request"
	"github.com/rlevidev/oauth-go/src/controllers/response"
	"github.com/rlevidev/oauth-go/src/models"
	"github.com/rlevidev/oauth-go/src/services"
)

func CreateUser(ctx *gin.Context) {
	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		ctx.JSON(restErr.Status, restErr)
		return
	}

	userDomain := models.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
	)

	userDomainResult, err := services.CreateUser(*userDomain)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	userResponse := response.UserResponse{
		ID:    userDomainResult.ID,
		Email: userDomainResult.Email,
		Name:  userDomainResult.Name,
	}

	ctx.JSON(201, userResponse)
}
