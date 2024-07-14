package service

import (
	"fmt"

	errorHandler "github.com/alexandreReys/api-jwt-repository-mongodb/src/configuration/error_handler"
	"github.com/alexandreReys/api-jwt-repository-mongodb/src/configuration/logger"
	model "github.com/alexandreReys/api-jwt-repository-mongodb/src/model/users"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *errorHandler.ErrorHandler {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println("User created: ", userDomain.GetPassword())

	return nil
}
