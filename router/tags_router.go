package router

import (
	"github.com/gin-gonic/gin"
	"keeper-crud/controller"
)

func NewTagsRouter(baseRouter *gin.RouterGroup) *gin.RouterGroup {
	return baseRouter.Group("/tags")
}

func SetupTagsRouter(baseRouter *gin.RouterGroup, tagsController *controller.TagsController) {
	tagsRouter := NewTagsRouter(baseRouter)
	tagsRouter.GET("", tagsController.FindAll)
	tagsRouter.GET("/:tagId", tagsController.FindById)
	tagsRouter.POST("", tagsController.Create)
	tagsRouter.PATCH("/:tagId", tagsController.Update)
	tagsRouter.DELETE("/:tagId", tagsController.Delete)
}
