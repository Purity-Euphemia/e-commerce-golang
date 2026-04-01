package models

import (
	"time"

	"gorm.io/gorm"
)

type Coupon struct {
	gorm.Model
	Code          string    `json:"code" gorm:"unique;not null"`
	DiscountType  string    `json:"discount_type" gorm:"default:percentage"` // percentage or fixed
	DiscountValue float64   `json:"discount_value"`
	MinAmount     float64   `json:"min_amount" gorm:"default:0"`
	MaxUsage      int       `json:"max_usage" gorm:"default:-1"` // -1 for unlimited
	UsageCount    int       `json:"usage_count" gorm:"default:0"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	IsActive      bool      `json:"is_active" gorm:"default:true"`
}
