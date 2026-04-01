package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ecommerce-go/services"
	"ecommerce-go/utils"
)

type CheckoutInput struct {
	Total float64 `json:"total" binding:"required,gt=0"`
}

func Checkout(c *gin.Context) {
	userID := c.GetUint("user_id")

	var input CheckoutInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	order, err := services.Checkout(userID, input.Total)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    order,
	})
}

func GetMyOrders(c *gin.Context) {
	userID := c.GetUint("user_id")

	orders, err := services.GetOrdersByUser(userID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, orders)
}
