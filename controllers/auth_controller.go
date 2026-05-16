package controllers

import (
	"net/http"

	"github.com/binojmadhuranga/BrainGrid-go/dto"
	"github.com/binojmadhuranga/BrainGrid-go/services"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var request dto.RegisterRequest

	// Validate request body
	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// Call service
	err := services.RegisterUser(request)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}

func Login(c *gin.Context) {

	var request dto.LoginRequest

	// Validate request body
	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// Call service
	token, user, err := services.LoginUser(
		request,
	)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}