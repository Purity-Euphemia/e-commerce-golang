package repositories

import (
	"ecommerce-go/database"
	"ecommerce-go/models"
)

func AddToCart(cartItem *models.CartItem) error {
	return database.DB.Create(cartItem).Error
}

func GetCartItemsByUser(userID uint) ([]models.CartItem, error) {
	var items []models.CartItem
	err := database.DB.Joins("JOIN carts ON carts.id = cart_items.cart_id").Where("carts.user_id = ?", userID).Find(&items).Error
	return items, err
}

func UpdateCartItem(userID, productID uint, qty int) error {
	return database.DB.
		Model(&models.CartItem{}).
		Where("product_id = ?", productID).
		Where("EXISTS(SELECT 1 FROM carts WHERE carts.id = cart_items.cart_id AND carts.user_id = ?)", userID).
		Update("quantity", qty).Error
}

func DeleteCartItem(userID, productID uint) error {
	return database.DB.
		Where("product_id = ?", productID).
		Where("EXISTS(SELECT 1 FROM carts WHERE carts.id = cart_items.cart_id AND carts.user_id = ?)", userID).
		Delete(&models.CartItem{}).Error
}
