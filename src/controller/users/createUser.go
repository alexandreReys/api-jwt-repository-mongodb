package controller

import (
	"net/http"

	"github.com/alexandreReys/api-jwt-repository-mongodb/src/configuration/logger"
	"github.com/alexandreReys/api-jwt-repository-mongodb/src/configuration/validation"
	"github.com/alexandreReys/api-jwt-repository-mongodb/src/controller/model/request"
	model "github.com/alexandreReys/api-jwt-repository-mongodb/src/model/users"
	"github.com/alexandreReys/api-jwt-repository-mongodb/src/model/users/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser Controller",
		zap.String("jorney", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user information", err,
			zap.String("jorney", "createUser"),
		)
		errorHandler := validation.ValidateUserError(err)
		c.JSON(errorHandler.Code, errorHandler)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email, 
		userRequest.Password, 
		userRequest.Name, 
		userRequest.Age,
	)
	service := service.NewUserDomainService()
	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created sucessfully",
		zap.String("jorney", "createUser"),
	)

	c.String(http.StatusOK, "")
}
	