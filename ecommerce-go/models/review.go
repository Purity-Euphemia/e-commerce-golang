package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ProductID uint     `json:"product_id"`
	Product   *Product `gorm:"foreignKey:ProductID"`
	UserID    uint     `json:"user_id"`
	User      *User    `gorm:"foreignKey:UserID"`
	Rating    int      `json:"rating" gorm:"check:rating >= 1 AND rating <= 5"`
	Title     string   `json:"title"`
	Comment   string   `json:"comment" gorm:"type:text"`
	Helpful   int      `json:"helpful" gorm:"default:0"`
}
