package controllers

import (
	"net/http"

	adminDto "github.com/binojmadhuranga/BrainGrid-go/dto/admin"
	"github.com/binojmadhuranga/BrainGrid-go/services"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	users, err := services.GetUsersService()

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

	err := services.UpdateUserRoleService(
		id,
		request.Role,
	)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Role updated successfully",
	})
}

func UpdateUserStatus(c *gin.Context) {

	id := c.Param("id")

	var request adminDto.UpdateStatusRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := services.UpdateUserStatusService(
		id,
		request.Status,
	)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Status updated successfully",
	})
}

func DeleteUser(c *gin.Context) {

	id := c.Param("id")

	err := services.DeleteUserService(id)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}