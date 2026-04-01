package models

import "gorm.io/gorm"

type Wishlist struct {
	gorm.Model
	UserID    uint     `json:"user_id"`
	User      *User    `gorm:"foreignKey:UserID"`
	ProductID uint     `json:"product_id"`
	Product   *Product `gorm:"foreignKey:ProductID"`
}
