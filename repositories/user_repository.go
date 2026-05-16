package repositories

import (
	"github.com/binojmadhuranga/BrainGrid-go/config"
	"github.com/binojmadhuranga/BrainGrid-go/models"
)

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func FindUserByEmail(email string) (models.User, error) {

	var user models.User

	err := config.DB.
		Where("email = ?", email).
		First(&user).Error

	return user, err
}