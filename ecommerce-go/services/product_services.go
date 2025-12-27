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
func ReduceStock(productID uint, qty int) error {
	var product models.Product
	err := config.DB.First(&product, productID).Error
	if err != nil {
		return err
	}

	if product.Stock < qty {
		return errors.New("not enough stock")
	}

	product.Stock -= qty
	return config.DB.Save(&product).Error
}


