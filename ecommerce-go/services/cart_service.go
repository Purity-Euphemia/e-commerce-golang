package services

import (
	"errors"

	"ecommerce-go/database"
	"ecommerce-go/models"
	"ecommerce-go/repositories"
)

func AddProductToCart(userID, productID uint, qty int) error {
	// Get or create cart for user
	var cart models.Cart
	if err := database.DB.Where("user_id = ?", userID).FirstOrCreate(&cart, models.Cart{UserID: userID}).Error; err != nil {
		return err
	}

	// Create cart item
	cartItem := models.CartItem{
		CartID:    cart.ID,
		ProductID: productID,
		Quantity:  qty,
	}
	return repositories.AddToCart(&cartItem)
}

func UpdateCart(userID, productID uint, qty int) error {
	if qty <= 0 {
		return errors.New("quantity must be greater than zero")
	}
	return repositories.UpdateCartItem(userID, productID, qty)
}

func RemoveFromCart(userID, productID uint) error {
	return repositories.DeleteCartItem(userID, productID)
}
