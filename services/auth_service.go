package services

import (
	"errors"

	"github.com/binojmadhuranga/BrainGrid-go/dto"
	"github.com/binojmadhuranga/BrainGrid-go/models"
	"github.com/binojmadhuranga/BrainGrid-go/repositories"
	"github.com/binojmadhuranga/BrainGrid-go/utils"
)

func RegisterUser(request dto.RegisterRequest) error {

	existingUser, _ := repositories.FindUserByEmail(
		request.Email,
	)

	// Check existing user
	if existingUser.ID != 0 {
		return errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(
		request.Password,
	)

	if err != nil {
		return err
	}

	// Create user model
	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
	}

	// Save user
	return repositories.CreateUser(&user)
}

func LoginUser(
	request dto.LoginRequest,
) (string, models.User, error) {

	user, err := repositories.FindUserByEmail(
		request.Email,
	)

	if err != nil {
		return "", models.User{}, errors.New(
			"invalid credentials",
		)
	}

	// Check password
	validPassword := utils.CheckPassword(
		request.Password,
		user.Password,
	)

	if !validPassword {
		return "", models.User{}, errors.New(
			"invalid credentials",
		)
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		return "", models.User{}, err
	}

	return token, user, nil
}