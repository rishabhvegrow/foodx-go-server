package repositories

import (
	"github.com/rishabhvegrow/foodx-go-server/models"
)

func GetUsers()(*[]models.User, error){
	var users []models.User
    if err := db.Find(&users).Error; err != nil {
        return nil, err
    }
	return &users, nil
}

func GetUser(userID any)(*models.User, error){
    var user models.User
    if err := db.First(&user, userID).Error; err != nil {
        return nil, err
    }

	return &user, nil
}

func CraeteUser(user models.User)(*models.User, error){
    if err := db.Create(&user).Error; err != nil {
        return nil, err
    }
	return &user, nil
}

func DeleteUser(userID any)(error){
	var user models.User
    if err:= db.Delete(&user, userID).Error; err != nil {
        return err
    }
	return nil
}