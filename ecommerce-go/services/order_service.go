package services

import (
	"ecommerce-go/models"
	"ecommerce-go/repositories"
)

func Checkout(userID uint, total float64) (*models.Order, error) {
	if err := ProcessPayment(total); err != nil {
		return nil, err
	}

	order := models.Order{
		UserID: userID,
		Total:  total,
	}
	err := repositories.CreateOrder(&order)
	return &order, err
}
