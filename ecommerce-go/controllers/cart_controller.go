package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ecommerce-go/services"
	"strconv"
)

type CartInput struct {
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

func AddToCart(c *gin.Context) {
	var input CartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err := services.AddProductToCart(input.UserID, input.ProductID, input.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product added to cart"})
}
func UpdateCartItem(c *gin.Context) {
	userID := c.GetUint("user_id")

	var input UpdateCartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err := services.UpdateCart(userID, input.ProductID, input.Quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "cart updated"})
}

func DeleteCartItem(c *gin.Context) {
	userID := c.GetUint("user_id")
	productIDParam := c.Param("product_id")

	productID64, _ := strconv.ParseUint(productIDParam, 10, 64)

	err := services.RemoveFromCart(userID, uint(productID64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "item removed"})
}
