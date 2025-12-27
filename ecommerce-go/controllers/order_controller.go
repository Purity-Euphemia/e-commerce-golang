package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ecommerce-go/services"
)

type CheckoutInput struct {
	Total float64 `json:"total"`
}

func Checkout(c *gin.Context) {
	userID := c.GetUint("user_id")

	var input CheckoutInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	order, err := services.Checkout(userID, input.Total)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, order)
}

func GetMyOrders(c *gin.Context) {
	userID := c.GetUint("user_id")

	orders, err := services.GetOrdersByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, orders)
}
