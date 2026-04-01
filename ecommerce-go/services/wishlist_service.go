package services

import (
	"ecommerce-go/models"
	"ecommerce-go/repositories"
)

func AddToWishlist(userID, productID uint) error {
	// Check if already in wishlist
	isInWishlist, _ := repositories.IsInWishlist(userID, productID)
	if isInWishlist {
		return repositories.RemoveFromWishlist(userID, productID)
	}
	return repositories.AddToWishlist(userID, productID)
}

func RemoveFromWishlist(userID, productID uint) error {
	return repositories.RemoveFromWishlist(userID, productID)
}

func GetUserWishlist(userID uint) ([]models.Wishlist, error) {
	return repositories.GetUserWishlist(userID)
}

func IsInWishlist(userID, productID uint) (bool, error) {
	return repositories.IsInWishlist(userID, productID)
}
