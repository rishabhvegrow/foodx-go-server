package repositories

import (
	"github.com/rishabhvegrow/foodx-go-server/database"
	"github.com/rishabhvegrow/foodx-go-server/models"
)

func GetRestaurents()(*[]models.Restaurant, error){
    db = database.GetDB()
	var restaurants []models.Restaurant
    if err := db.Find(&restaurants).Error; err != nil {
        return nil, err
    }
	return &restaurants, nil
}

func GetRestaurent(restID any)(*models.Restaurant, error){
    db = database.GetDB()
    var restaurent models.Restaurant
    if err := db.First(&restaurent, restID).Error; err != nil {
        return nil, err
    }

	return &restaurent, nil
}

func CraeteRestaurent(restaurent models.Restaurant)(*models.Restaurant, error){
    db = database.GetDB()
    if err := db.Create(&restaurent).Error; err != nil {
        return nil, err
    }
	return &restaurent, nil
}

func DeleteRestaurent(restID any)(error){
    db = database.GetDB()
	var restaurent models.Restaurant
    if err:= db.Delete(&restaurent, restID).Error; err != nil {
        return err
    }
	return nil
}