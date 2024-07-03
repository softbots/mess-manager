package repositories

import (
	"mess-manager/pkg/domain"
	"mess-manager/pkg/models"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func UserDbInstance(db *gorm.DB) domain.IUserRepo {
	return &userRepo{
		db: db,
	}
}

func (repo *userRepo) SignUp(user *models.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}
