package service

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"keeper-crud/data/request"
	"keeper-crud/helper"
	"keeper-crud/model"
	"keeper-crud/repository"
)

type UsersServiceImplementation struct {
	UsersRepository repository.UsersRepository
	Validate        *validator.Validate
}

func NewUsersServiceImplementation(userRepository repository.UsersRepository, validate *validator.Validate) UsersService {
	return &UsersServiceImplementation{
		UsersRepository: userRepository,
		Validate:        validate,
	}
}

func (u *UsersServiceImplementation) SignUp(user request.UserSignUpRequest) {
	err := u.Validate.Struct(user)
	helper.ErrorPanic(err)

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	helper.ErrorPanic(err)
	hashedPassword := string(passwordBytes)

	userModel := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	}
	u.UsersRepository.SignUp(userModel)
}
