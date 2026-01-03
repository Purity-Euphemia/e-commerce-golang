package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID          uint
	ProductID        uint
	Product          Product
	Quantity         int
	PriceAtPurchase float64
}
