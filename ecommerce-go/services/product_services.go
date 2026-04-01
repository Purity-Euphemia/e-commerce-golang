package services

import (
	"errors"

	"ecommerce-go/database"
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
		Stock: 0,
	}
	err := repositories.CreateProduct(&product)
	return &product, err
}

func ReduceStock(productID uint, qty int) error {
	var product models.Product
	err := database.DB.First(&product, productID).Error
	if err != nil {
		return err
	}

	if product.Stock < qty {
		return errors.New("not enough stock")
	}

	product.Stock -= qty
	return database.DB.Save(&product).Error
}
