package service

import (
	errorHandler "github.com/alexandreReys/api-jwt-repository-mongodb/src/configuration/error_handler"
	model "github.com/alexandreReys/api-jwt-repository-mongodb/src/model/users"
)

func (ud *userDomainService) FindUser(string) (
	*model.UserDomainInterface, *errorHandler.ErrorHandler,
) {
	return nil, nil
}
