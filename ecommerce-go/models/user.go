package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string     `json:"name" gorm:"not null"`
	Email    string     `gorm:"unique;not null" json:"email"`
	Password string     `json:"-"`
	Phone    string     `json:"phone"`
	Role     string     `json:"role" gorm:"default:customer"`
	Avatar   string     `json:"avatar"`
	Address  string     `json:"address" gorm:"type:text"`
	City     string     `json:"city"`
	State    string     `json:"state"`
	ZipCode  string     `json:"zip_code"`
	IsActive bool       `json:"is_active" gorm:"default:true"`
	Cart     *Cart      `gorm:"foreignKey:UserID"`
	Orders   []Order    `gorm:"foreignKey:UserID"`
	Reviews  []Review   `gorm:"foreignKey:UserID"`
	Wishlist []Wishlist `gorm:"foreignKey:UserID"`
}
