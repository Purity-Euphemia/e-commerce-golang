package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ecommerce-go/services"
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
