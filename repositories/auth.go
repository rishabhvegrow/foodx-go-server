package repositories

import (
    "gorm.io/gorm"
	"github.com/rishabhvegrow/foodx-go-server/database"
	"github.com/rishabhvegrow/foodx-go-server/models"
)


var db *gorm.DB

func GetUserByEmail(email string) (*models.User, error) {
	db = database.GetDB()
    var user models.User
    if err := db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}