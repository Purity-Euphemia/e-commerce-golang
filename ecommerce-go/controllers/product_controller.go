package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ecommerce-go/services"
	"ecommerce-go/utils"
)

type ProductInput struct {
	Name          string  `json:"name" binding:"required"`
	Description   string  `json:"description" binding:"required"`
	Price         float64 `json:"price" binding:"required,gt=0"`
	DiscountPrice float64 `json:"discount_price" binding:"gte=0"`
	Stock         int     `json:"stock" binding:"required,gte=0"`
	CategoryID    uint    `json:"category_id" binding:"required"`
	Image         string  `json:"image"`
	SKU           string  `json:"sku"`
	Slug          string  `json:"slug" binding:"required"`
}

func GetProducts(c *gin.Context) {
	page, pageSize := utils.GetPaginationFromQuery(c)

	products, total, err := services.GetProductsPaginated(page, pageSize)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    products,
		"pagination": gin.H{
			"page":      page,
			"page_size": pageSize,
			"total":     total,
		},
	})
}

func SearchProducts(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		utils.Error(c, http.StatusBadRequest, "search query required")
		return
	}

	page, pageSize := utils.GetPaginationFromQuery(c)

	products, total, err := services.SearchProducts(query, page, pageSize)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    products,
		"pagination": gin.H{
			"page":      page,
			"page_size": pageSize,
			"total":     total,
		},
	})
}

func GetProductsByCategory(c *gin.Context) {
	categoryIDStr := c.Param("category_id")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid category id")
		return
	}

	page, pageSize := utils.GetPaginationFromQuery(c)

	products, total, err := services.GetProductsByCategory(uint(categoryID), page, pageSize)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    products,
		"pagination": gin.H{
			"page":      page,
			"page_size": pageSize,
			"total":     total,
		},
	})
}

func GetProductByID(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid product id")
		return
	}

	product, err := services.GetProductByID(uint(productID))
	if err != nil {
		utils.Error(c, http.StatusNotFound, "product not found")
		return
	}

	utils.Success(c, product)
}

func CreateProduct(c *gin.Context) {
	var input ProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	product, err := services.CreateProduct(input.Name, input.Description, input.Price, input.DiscountPrice, input.Stock, input.CategoryID, input.Slug, input.Image, input.SKU)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    product,
	})
}

func UpdateProduct(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid product id")
		return
	}

	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	product, err := services.UpdateProduct(uint(productID), input.Name, input.Description, input.Price, input.DiscountPrice, input.Stock, input.CategoryID, input.Slug, input.Image, input.SKU)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, product)
}

func DeleteProduct(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid product id")
		return
	}

	err = services.DeleteProduct(uint(productID))
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, gin.H{"message": "product deleted"})
}
