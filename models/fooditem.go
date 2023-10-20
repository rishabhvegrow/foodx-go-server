package models

import (
	"gorm.io/gorm"
)

type FoodItem struct {
    gorm.Model
	Name        string  `gorm:"not null"`
	Price       float32 `gorm:"not null"`
	Description string
    RestaurantID uint    `gorm:"not null"`
}