package models

import (
	"gorm.io/gorm"
)

type CartItem struct {
    gorm.Model
    UserID       uint   `gorm:"not null"`
    FoodItemID   uint   `gorm:"not null"`
    Quantity     uint
    Price        float32 `gorm:"not null"`
    TransactionID uint
    IsCheckedOut bool `gorm:"default:false"`
    FoodItem     FoodItem
}