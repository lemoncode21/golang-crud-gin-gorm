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

//	@title			Tag Service API
//	@version		1.0
//	@description	A Tag service API in Go using Gin framework

// @host		localhost:8888
// @BasePath	/api
func main() {

	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	// Repository
	tagsRepository := repository.NewTagsREpositoryImpl(db)

	// Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	// Controller
	tagsController := controller.NewTagsController(tagsService)

	// Router
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
