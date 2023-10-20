package models

import (
	"gorm.io/gorm"
)

type Restaurant struct {
    gorm.Model
    Name     string `gorm:"not null"`
    Type     string `gorm:"not null"`
    Address  string
    Contact  string `gorm:"not null"`
    Rating   float32
    ImageUrl string `gorm:"not null"`
}