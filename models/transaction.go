package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
    gorm.Model
    UserID uint	  `gorm:"not null"`
    Total  float64
}