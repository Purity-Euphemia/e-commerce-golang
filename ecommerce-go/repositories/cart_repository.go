package repositories

import (
	"ecommerce-go/config"
	"ecommerce-go/models"
)

func AddToCart(cart *models.Cart) error {
	return config.DB.Create(cart).Error
}
