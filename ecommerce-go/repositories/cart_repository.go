package repositories

import (
	"ecommerce-go/config"
	"ecommerce-go/models"
)

func AddToCart(cart *models.Cart) error {
	return config.DB.Create(cart).Error
}
func GetCartItemsByUser(userID uint) ([]models.Cart, error) {
	var items []models.Cart
	err := config.DB.Where("user_id = ?", userID).Find(&items).Error
	return items, err
}

