package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID    uint     `json:"cart_id"`
	Cart      *Cart    `gorm:"foreignKey:CartID"`
	ProductID uint     `json:"product_id"`
	Product   *Product `gorm:"foreignKey:ProductID"`
	Quantity  int      `json:"quantity"`
}
