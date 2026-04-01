package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ecommerce-go/services"
	"ecommerce-go/utils"
)

type CartInput struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

type UpdateCartInput struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

func AddToCart(c *gin.Context) {
	userID := c.GetUint("user_id")
	var input CartInput

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err := services.AddProductToCart(userID, input.ProductID, input.Quantity)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, gin.H{"message": "Product added to cart"})
}

func UpdateCartItem(c *gin.Context) {
	userID := c.GetUint("user_id")

	var input UpdateCartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err := services.UpdateCart(userID, input.ProductID, input.Quantity)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, gin.H{"message": "cart updated"})
}

func DeleteCartItem(c *gin.Context) {
	userID := c.GetUint("user_id")
	productIDParam := c.Param("product_id")

	productID64, err := strconv.ParseUint(productIDParam, 10, 64)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid product id")
		return
	}

	err = services.RemoveFromCart(userID, uint(productID64))
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, gin.H{"message": "item removed"})
}
