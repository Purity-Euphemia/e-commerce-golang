package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ecommerce-go/services"
	"ecommerce-go/utils"
)

type ReviewInput struct {
	Rating  int    `json:"rating" binding:"required,min=1,max=5"`
	Title   string `json:"title" binding:"required"`
	Comment string `json:"comment" binding:"required"`
}

func AddReview(c *gin.Context) {
	userID := c.GetUint("user_id")
	productIDStr := c.Param("product_id")
	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid product id")
		return
	}

	var input ReviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	review, err := services.AddReview(uint(productID), userID, input.Rating, input.Title, input.Comment)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// Update product rating
	services.UpdateProductRating(uint(productID))

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    review,
	})
}

func GetProductReviews(c *gin.Context) {
	productIDStr := c.Param("product_id")
	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid product id")
		return
	}

	page, pageSize := utils.GetPaginationFromQuery(c)

	reviews, total, err := services.GetProductReviews(uint(productID), page, pageSize)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    reviews,
		"pagination": gin.H{
			"page":      page,
			"page_size": pageSize,
			"total":     total,
		},
	})
}

func DeleteReview(c *gin.Context) {
	userID := c.GetUint("user_id")
	reviewIDStr := c.Param("review_id")
	reviewID, err := strconv.ParseUint(reviewIDStr, 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid review id")
		return
	}

	err = services.DeleteReview(uint(reviewID), userID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, gin.H{"message": "review deleted"})
}
