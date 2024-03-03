package service

import (
	"keeper-crud/data/request"
	"keeper-crud/model"
)

type UsersService interface {
	SignUp(user request.UserSignUpRequest)
	AuthenticateUser(email string, password string) (*model.User, error)
}
