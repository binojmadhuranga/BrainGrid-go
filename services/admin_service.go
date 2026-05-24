package services

import (
    "github.com/binojmadhuranga/BrainGrid-go/repositories"
)

func GetUsersService() interface{} {

    users, err := repositories.GetAllUsers()

    if err != nil {
        return nil
    }

    return users
}