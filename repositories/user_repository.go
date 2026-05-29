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

func FindUserByID(id string) (models.User, error) {

	var user models.User

	err := config.DB.
		First(&user, id).Error

	return user, err
}

func GetAllUsers() ([]models.User, error) {

	var users []models.User

	err := config.DB.
		Where("role != ?", "ADMIN").
		Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func UpdateUser(user *models.User) error {

	return config.DB.Save(user).Error
}

func DeleteUser(id string) error {

	return config.DB.Delete(
		&models.User{},
		id,
	).Error
}