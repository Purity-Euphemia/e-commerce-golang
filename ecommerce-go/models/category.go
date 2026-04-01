package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string    `json:"name" gorm:"unique;not null"`
	Description string    `json:"description"`
	Slug        string    `json:"slug" gorm:"unique;not null"`
	Icon        string    `json:"icon"`
	Products    []Product `gorm:"foreignKey:CategoryID"`
}
