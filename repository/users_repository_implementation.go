package repository

import (
	"gorm.io/gorm"
	"keeper-crud/helper"
	"keeper-crud/model"
)

type UsersRepositoryImplementation struct {
	Db *gorm.DB
}

func NewUsersRepositoryImplementation(Db *gorm.DB) UsersRepository {
	return &UsersRepositoryImplementation{Db: Db}
}

func (u UsersRepositoryImplementation) SignUp(user model.User) {
	result := u.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}
