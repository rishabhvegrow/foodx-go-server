package repositories

import (
	"github.com/rishabhvegrow/foodx-go-server/models"
)

func GetRestaurents()(*[]models.Restaurant, error){
	var restaurants []models.Restaurant
    if err := db.Find(&restaurants).Error; err != nil {
        return nil, err
    }
	return &restaurants, nil
}

func GetRestaurent(restID any)(*models.Restaurant, error){
    var restaurent models.Restaurant
    if err := db.First(&restaurent, restID).Error; err != nil {
        return nil, err
    }

	return &restaurent, nil
}

func CraeteRestaurent(restaurent models.Restaurant)(*models.Restaurant, error){
    if err := db.Create(&restaurent).Error; err != nil {
        return nil, err
    }
	return &restaurent, nil
}

func DeleteRestaurent(restID any)(error){
	var restaurent models.Restaurant
    if err:= db.Delete(&restaurent, restID).Error; err != nil {
        return err
    }
	return nil
}