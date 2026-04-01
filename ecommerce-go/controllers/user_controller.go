package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ecommerce-go/services"
	"ecommerce-go/utils"
)

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := services.RegisterUser(input.Name, input.Email, input.Password)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, user)
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := services.LoginUser(input.Email, input.Password)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.Success(c, user)
}
