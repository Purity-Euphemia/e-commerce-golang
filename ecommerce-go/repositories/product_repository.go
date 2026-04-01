package repositories

import (
	"ecommerce-go/database"
	"ecommerce-go/models"
)

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	result := database.DB.Find(&products)
	return products, result.Error
}

func CreateProduct(product *models.Product) error {
	result := database.DB.Create(product)
	return result.Error
}
