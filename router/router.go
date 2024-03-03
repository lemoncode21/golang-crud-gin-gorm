package router

import (
	"keeper-crud/controller"

	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(tagsController *controller.TagsController, usersController *controller.UsersController) *gin.Engine {
	router := gin.Default()

	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	SetupTagsRouter(baseRouter, tagsController)
	SetupUsersRouter(baseRouter, usersController)

	return router
}
