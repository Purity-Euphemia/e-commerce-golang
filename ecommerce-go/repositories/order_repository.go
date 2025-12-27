package repositories

import (
	"ecommerce-go/config"
	"ecommerce-go/models"
)

func CreateOrder(order *models.Order) error {
	return config.DB.Create(order).Error
}

func GetOrdersByUser(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := config.DB.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}
