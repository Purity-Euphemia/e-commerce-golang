package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID      uint        `json:"user_id"`
	User        *User       `gorm:"foreignKey:UserID"`
	TotalAmount float64     `json:"total_amount"`
	Status      string      `json:"status" gorm:"default:pending"`
	Items       []OrderItem `gorm:"foreignKey:OrderID"`
}
