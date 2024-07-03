package domain

import "mess-manager/pkg/models"

type IUserRepo interface {
	SignUp(user *models.User) error
}

type IUserService interface {
	SignUp(user *models.User) error
}
