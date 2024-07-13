package controller

import (
	"net/http"

	"github.com/alexandreReys/api-jwt-repository-mongodb/src/configuration/logger"
	"github.com/alexandreReys/api-jwt-repository-mongodb/src/configuration/validation"
	"github.com/alexandreReys/api-jwt-repository-mongodb/src/controller/model/request"
	"github.com/alexandreReys/api-jwt-repository-mongodb/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Creating a new user",
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

	response := response.UserResponse{
		Id:       "test",
		Email:    userRequest.Email,
		Name:     userRequest.Name,
		Age: 			userRequest.Age,
	}

	logger.Info("User created sucessfully",
		zap.String("jorney", "createUser"),
	)

	c.JSON(http.StatusCreated, response)
}
	