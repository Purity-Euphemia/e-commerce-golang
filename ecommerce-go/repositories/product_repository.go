package repositories

import (
	"ecommerce-go/config"
	"ecommerce-go/models"
)

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	result := config.DB.Find(&products)
	return products, result.Error
}

func CreateProduct(product *models.Product) error {
	result := config.DB.Create(product)
	return result.Error
}
