package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ecommerce-go/services"
	"ecommerce-go/utils"
)

type RegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := services.RegisterUser(input.Name, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := services.LoginUser(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
