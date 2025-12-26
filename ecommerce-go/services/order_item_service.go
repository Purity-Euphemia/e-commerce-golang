package services

import (
	"ecommerce-go/models"
	"ecommerce-go/repositories"
)

func AddOrderItem(orderID, productID uint, qty int, price float64) error {
	item := models.OrderItem{
		OrderID:   orderID,
		ProductID: productID,
		Quantity:  qty,
		Price:     price,
	}
	return repositories.CreateOrderItem(&item)
}
