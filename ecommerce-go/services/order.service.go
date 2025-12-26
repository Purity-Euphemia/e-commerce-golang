package services

import (
	"ecommerce-go/models"
	"ecommerce-go/repositories"
)

func Checkout(userID uint, total float64) (*models.Order, error) {
	order := models.Order{
		UserID: userID,
		Total:  total,
	}
	err := repositories.CreateOrder(&order)
	return &order, err
}
