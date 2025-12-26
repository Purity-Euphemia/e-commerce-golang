package repositories

import (
	"ecommerce-go/config"
	"ecommerce-go/models"
)

func CreateOrder(order *models.Order) error {
	return config.DB.Create(order).Error
}
