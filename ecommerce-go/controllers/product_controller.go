package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ecommerce-go/services"
	"ecommerce-go/utils"
)

type ProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func GetProducts(c *gin.Context) {
	products, err := services.GetProducts()
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, products)
}

func CreateProduct(c *gin.Context) {
	var input ProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	product, err := services.CreateProduct(input.Name, input.Price)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, product)
}
