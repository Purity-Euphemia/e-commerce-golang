package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string    `json:"name" gorm:"not null;index"`
	Slug          string    `json:"slug" gorm:"unique;not null"`
	Description   string    `json:"description" gorm:"type:text"`
	Price         float64   `json:"price" gorm:"not null"`
	DiscountPrice float64   `json:"discount_price" gorm:"default:0"`
	Stock         int       `json:"stock" gorm:"not null"`
	CategoryID    uint      `json:"category_id"`
	Category      *Category `gorm:"foreignKey:CategoryID"`
	Image         string    `json:"image"`
	Images        string    `json:"images" gorm:"type:text"` // JSON array of image URLs
	SKU           string    `json:"sku" gorm:"unique"`
	Rating        float64   `json:"rating" gorm:"default:0"`
	Reviews       []Review  `gorm:"foreignKey:ProductID"`
	CreatedAt     int64     `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt     int64     `json:"updated_at" gorm:"autoUpdateTime:milli"`
}
