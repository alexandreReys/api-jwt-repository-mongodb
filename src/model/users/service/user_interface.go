package service

import (
	errorHandler "github.com/alexandreReys/api-jwt-repository-mongodb/src/configuration/error_handler"
	model "github.com/alexandreReys/api-jwt-repository-mongodb/src/model/users"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *errorHandler.ErrorHandler
	UpdateUser(string, model.UserDomainInterface) *errorHandler.ErrorHandler
	FindUser(string) (*model.UserDomainInterface, *errorHandler.ErrorHandler)
	DeleteUser(string) *errorHandler.ErrorHandler
}
