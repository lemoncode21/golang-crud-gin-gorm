package main

import (
	"keeper-crud/config"
	"keeper-crud/controller"
	_ "keeper-crud/docs"
	"keeper-crud/helper"
	"keeper-crud/model"
	"keeper-crud/repository"
	"keeper-crud/router"
	"keeper-crud/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

//	@title			Keeper API
//	@version		1.0
//	@description	A Keeper API in Go using Gin framework

// @host		localhost:8888
// @BasePath	/api
func main() {

	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	err := db.Table("tags").AutoMigrate(&model.Tags{})
	helper.ErrorPanic(err)
	err = db.Table("users").AutoMigrate(&model.User{})
	helper.ErrorPanic(err)

	// Repository
	tagsRepository := repository.NewTagsREpositoryImpl(db)
	usersRepository := repository.NewUsersRepositoryImplementation(db)

	// Service
	mainService := service.NewService(tagsRepository, usersRepository, validate)

	// Controllers
	mainController := controller.NewController(&mainService.TagsService, &mainService.UsersService)

	// Router
	routes := router.NewRouter(mainController.TagsController, mainController.UsersController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err = server.ListenAndServe()
	helper.ErrorPanic(err)
}
