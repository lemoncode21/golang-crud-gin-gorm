package repository

import "keeper-crud/model"

type UsersRepository interface {
	SignUp(user model.User)
}
