package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ecommerce-go/services"
	"ecommerce-go/utils"
)

type WishlistInput struct {
	ProductID uint `json:"product_id" binding:"required"`
}

func AddToWishlist(c *gin.Context) {
	userID := c.GetUint("user_id")

	var input WishlistInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err := services.AddToWishlist(userID, input.ProductID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, gin.H{"message": "added to wishlist"})
}

func RemoveFromWishlist(c *gin.Context) {
	userID := c.GetUint("user_id")
	productIDStr := c.Param("product_id")

	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid product id")
		return
	}

	err = services.RemoveFromWishlist(userID, uint(productID))
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, gin.H{"message": "removed from wishlist"})
}

func GetMyWishlist(c *gin.Context) {
	userID := c.GetUint("user_id")

	wishlist, err := services.GetUserWishlist(userID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, wishlist)
}
