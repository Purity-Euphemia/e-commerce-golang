package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID         uint     `json:"order_id"`
	Order           *Order   `gorm:"foreignKey:OrderID"`
	ProductID       uint     `json:"product_id"`
	Product         *Product `gorm:"foreignKey:ProductID"`
	Quantity        int      `json:"quantity"`
	PriceAtPurchase float64  `json:"price_at_purchase"`
}
