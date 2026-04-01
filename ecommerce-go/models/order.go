package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID          uint        `json:"user_id"`
	User            *User       `gorm:"foreignKey:UserID"`
	OrderNumber     string      `json:"order_number" gorm:"unique"`
	TotalAmount     float64     `json:"total_amount"`
	DiscountAmount  float64     `json:"discount_amount" gorm:"default:0"`
	CouponCode      string      `json:"coupon_code"`
	Coupon          *Coupon     `gorm:"foreignKey:CouponCode;references:Code"`
	Status          string      `json:"status" gorm:"default:pending;index"`   // pending, confirmed, shipped, delivered, cancelled
	PaymentStatus   string      `json:"payment_status" gorm:"default:pending"` // pending, completed, failed
	ShippingAddress string      `json:"shipping_address"`
	Items           []OrderItem `gorm:"foreignKey:OrderID"`
	TrackingNumber  string      `json:"tracking_number"`
	Notes           string      `json:"notes" gorm:"type:text"`
}
