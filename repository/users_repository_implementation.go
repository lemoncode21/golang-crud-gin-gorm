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

func (u UsersRepositoryImplementation) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.Db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
