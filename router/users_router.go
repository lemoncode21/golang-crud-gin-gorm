package router

import (
	"github.com/gin-gonic/gin"
	"keeper-crud/controller"
)

func NewUsersRouter(baseRouter *gin.RouterGroup) *gin.RouterGroup {
	return baseRouter.Group("/users")
}

func SetupUsersRouter(baseRouter *gin.RouterGroup, usersController *controller.UsersController) {
	usersRouter := NewUsersRouter(baseRouter)
	usersRouter.POST("/signup", usersController.Signup)
}
