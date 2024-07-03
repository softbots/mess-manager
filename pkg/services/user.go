package services

import (
	"errors"
	"mess-manager/pkg/domain"
	"mess-manager/pkg/models"
)

type UserService struct {
	repo domain.IUserRepo
}

func UserServiceInstance(userRepo domain.IUserRepo) domain.IUserService {
	return &UserService{
		repo: userRepo,
	}
}

func (service *UserService) SignUp(user *models.User) error {
	if err := service.repo.SignUp(user); err != nil {
		return errors.New("User was not created")
	}

	return nil
}
