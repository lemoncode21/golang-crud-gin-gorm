package service

import "keeper-crud/data/request"

type UsersService interface {
	SignUp(user request.UserSignUpRequest)
}
