package service

import (
	"github.com/go-playground/validator/v10"
	"keeper-crud/repository"
)

type Service struct {
	TagsService  TagsService
	UsersService UsersService
}

func NewService(tagsRepository repository.TagsRepository, usersRepository repository.UsersRepository,
	validate *validator.Validate) *Service {
	return &Service{
		TagsService:  NewTagsServiceImpl(tagsRepository, validate),
		UsersService: NewUsersServiceImplementation(usersRepository, validate),
	}
}
