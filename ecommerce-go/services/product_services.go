package services

import (
	"ecommerce-go/models"
	"ecommerce-go/repositories"
)

func GetProducts() ([]models.Product, error) {
	return repositories.GetAllProducts()
}

func CreateProduct(name string, price float64) (*models.Product, error) {
	product := models.Product{
		Name:  name,
		Price: price,
	}
	err := repositories.CreateProduct(&product)
	return &product, err
}
