package services

import (
	"errors"

	"github.com/binojmadhuranga/BrainGrid-go/models"
	"github.com/binojmadhuranga/BrainGrid-go/repositories"
)

func GetUsersService() ([]models.User, error) {

	users, err := repositories.GetAllUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func UpdateUserRoleService(
	id string,
	role string,
) error {

	user, err := repositories.FindUserByID(id)

	if err != nil {
		return errors.New("user not found")
	}

	user.Role = role

	err = repositories.UpdateUser(&user)

	if err != nil {
		return err
	}

	return nil
}

func UpdateUserStatusService(
	id string,
	status string,
) error {

	user, err := repositories.FindUserByID(id)

	if err != nil {
		return errors.New("user not found")
	}

	user.Status = status

	err = repositories.UpdateUser(&user)

	if err != nil {
		return err
	}

	return nil
}

func DeleteUserService(id string) error {

	_, err := repositories.FindUserByID(id)

	if err != nil {
		return errors.New("user not found")
	}

	err = repositories.DeleteUser(id)

	if err != nil {
		return err
	}

	return nil
}