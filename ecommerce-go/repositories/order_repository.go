package repositories

import (
	"ecommerce-go/database"
	"ecommerce-go/models"
)

func CreateOrder(order *models.Order) error {
	return database.DB.Create(order).Error
}

func GetOrdersByUser(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := database.DB.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}
