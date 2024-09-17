package service

import (
	"github.com/oliverreese/golang-dev/pkg/model"
	"github.com/oliverreese/golang-dev/pkg/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) UserService {
	return UserService{
		Repo: repo,
	}
}

func (us *UserService) GetOneUserById(id int) (*model.User, error) {
	user, err := us.Repo.FindOneById(id)
	return user, err
}
