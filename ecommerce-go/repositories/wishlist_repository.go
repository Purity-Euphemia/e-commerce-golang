package repositories

import (
	"ecommerce-go/database"
	"ecommerce-go/models"
)

func AddToWishlist(userID, productID uint) error {
	return database.DB.Create(&models.Wishlist{UserID: userID, ProductID: productID}).Error
}

func RemoveFromWishlist(userID, productID uint) error {
	return database.DB.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&models.Wishlist{}).Error
}

func GetUserWishlist(userID uint) ([]models.Wishlist, error) {
	var wishlist []models.Wishlist
	err := database.DB.Where("user_id = ?", userID).Preload("Product").Find(&wishlist).Error
	return wishlist, err
}

func IsInWishlist(userID, productID uint) (bool, error) {
	var count int64
	err := database.DB.Model(&models.Wishlist{}).
		Where("user_id = ? AND product_id = ?", userID, productID).
		Count(&count).Error
	return count > 0, err
}
