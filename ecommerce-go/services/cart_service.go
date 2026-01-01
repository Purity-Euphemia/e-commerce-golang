package services

import "ecommerce-go/models"
import "ecommerce-go/repositories"

func AddProductToCart(userID, productID uint, qty int) error {
	cart := models.Cart{
		UserID:    userID,
		ProductID: productID,
		Quantity:  qty,
	}
	return repositories.AddToCart(&cart)
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

