package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"keeper-crud/data/request"
	"keeper-crud/data/response"
	"keeper-crud/helper"
	"keeper-crud/service"
	"net/http"
)

type UsersController struct {
	usersService service.UsersService
}

func NewUsersController(service service.UsersService) *UsersController {
	return &UsersController{usersService: service}
}

// Signup SignUp Users		godoc
//
//	@Summary		SignUp users
//	@Description	Save users data in Db.
//	@Param			users	body	request.UserSignUpRequest	true	"Signup users"
//	@Produce		application/json
//	@Users			users
//	@Success		200	{object}	response.Response{}
//	@Router			/signup [post]
func (controller *UsersController) Signup(ctx *gin.Context) {
	log.Info().Msg("signup user")
	userSignUpRequest := request.UserSignUpRequest{}
	err := ctx.ShouldBindJSON(&userSignUpRequest)
	helper.ErrorPanic(err)

	controller.usersService.SignUp(userSignUpRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
