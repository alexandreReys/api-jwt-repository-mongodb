package routes

import (
	controller "github.com/alexandreReys/api-jwt-repository-mongodb/src/controller/users"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/users/getUserById:userId", controller.FindUserById)
	r.GET("/users/getUserByEmail:userEmail", controller.FindUserByEmail)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users:userId", controller.UpdateUser)
	r.DELETE("/users:userId", controller.DeleteUser)
	
}
