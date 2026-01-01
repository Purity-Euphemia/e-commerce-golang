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
func UpdateCartItem(userID, productID uint, qty int) error {
	return config.DB.
		Model(&models.Cart{}).
		Where("user_id = ? AND product_id = ?", userID, productID).
		Update("quantity", qty).Error
}

func DeleteCartItem(userID, productID uint) error {
	return config.DB.
		Where("user_id = ? AND product_id = ?", userID, productID).
		Delete(&models.Cart{}).Error
}


