package controllers

import (
	"net/http"

	adminDto "github.com/binojmadhuranga/BrainGrid-go/dto/admin"
	"github.com/binojmadhuranga/BrainGrid-go/repositories"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	users, err := repositories.GetAllUsers()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch users",
		})

		return
	}

	c.JSON(http.StatusOK, users)
}

func UpdateUserRole(c *gin.Context) {

	id := c.Param("id")

	var request adminDto.UpdateRoleRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	user, err := repositories.FindUserByID(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})

		return
	}

	user.Role = request.Role

	repositories.UpdateUser(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Role updated",
	})
}

func UpdateUserStatus(c *gin.Context) {

	id := c.Param("id")

	var request adminDto.UpdateStatusRequest

	c.ShouldBindJSON(&request)

	user, err := repositories.FindUserByID(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})

		return
	}

	user.Status = request.Status

	repositories.UpdateUser(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Status updated",
	})
}

func DeleteUser(c *gin.Context) {

	id := c.Param("id")

	err := repositories.DeleteUser(id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Delete failed",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})
}
