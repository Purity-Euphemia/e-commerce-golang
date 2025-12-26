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
