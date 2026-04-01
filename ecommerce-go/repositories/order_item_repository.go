package repositories

import (
	"ecommerce-go/database"
	"ecommerce-go/models"
)

func CreateOrderItem(item *models.OrderItem) error {
	return database.DB.Create(item).Error
}
