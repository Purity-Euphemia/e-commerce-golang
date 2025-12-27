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
func ReduceStock(productID, qty int) error {
	for i, p := range products {
		if p.ID == productID {
			if p.Stock < qty {
				return errors.New("not enough stock")
			}
			products[i].Stock -= qty
			return nil
		}
	}
	return errors.New("product not found")
}

