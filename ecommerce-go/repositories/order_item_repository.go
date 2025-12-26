package repositories

import (
	"ecommerce-go/config"
	"ecommerce-go/models"
)

func CreateOrderItem(item *models.OrderItem) error {
	return config.DB.Create(item).Error
}
