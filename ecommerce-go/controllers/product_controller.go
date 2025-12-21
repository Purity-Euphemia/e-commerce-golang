package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"ecommerce-go/services"
)

func GetProducts(c *gin.Context) {
	products := services.GetAllProducts()
	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	var product services.ProductInput

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProduct := services.CreateProduct(product)
	c.JSON(http.StatusCreated, newProduct)
}
