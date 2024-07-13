package controller

import (
	"fmt"

	"github.com/alexandreReys/api-jwt-repository-mongodb/src/configuration/validation"
	"github.com/alexandreReys/api-jwt-repository-mongodb/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errorHandler := validation.ValidateUserError(err)

		c.JSON(errorHandler.Code, errorHandler)
		return
	}

	fmt.Println(userRequest)


}
	