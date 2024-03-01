package controller

import "keeper-crud/service"

type Controller struct {
	TagsController  *TagsController
	UsersController *UsersController
}

func NewController(tagsService *service.TagsService, usersService *service.UsersService) *Controller {
	return &Controller{
		TagsController:  NewTagsController(*tagsService),
		UsersController: NewUsersController(*usersService),
	}
}
