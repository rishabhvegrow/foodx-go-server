package repositories

import (
    "github.com/rishabhvegrow/foodx-go-server/database"
    "github.com/rishabhvegrow/foodx-go-server/models"
)


func CreateFoodItem(foodItem models.FoodItem)(*models.FoodItem,error){
	db = database.GetDB()
    if err := db.Create(&foodItem).Error; err != nil{
        return nil,err
    }
	return &foodItem, nil
}

func GetFoodItemByID(id any)(*models.FoodItem,error){
    db = database.GetDB()
	var foodItem models.FoodItem

    if err := db.First(&foodItem, id).Error; err != nil {
        return nil, err
    }
	return &foodItem, nil
}

func UpdateFoodItem(id any,updatedFoodItem models.FoodItem)(*models.FoodItem, error){
    db = database.GetDB()
	foodItem, err := GetFoodItemByID(id)
    if err != nil {
        return nil,err
    }

	foodItem.Name = updatedFoodItem.Name
    foodItem.Price = updatedFoodItem.Price
    foodItem.Description = updatedFoodItem.Description
    foodItem.RestaurantID = updatedFoodItem.RestaurantID
    foodItem.ImageUrl = updatedFoodItem.ImageUrl

    if err:=db.Save(&foodItem).Error; err != nil {
        return nil, err
    }

	return foodItem, nil
}

func DeleteFoodItem(id any)(error){
    db = database.GetDB()
    var foodItem models.FoodItem
    if err := db.Delete(&foodItem, id).Error; err != nil {
        return err
    }
    return nil
}

func GetMenuOfRestaurent(restaurentID any)(*[]models.FoodItem, error){
    db = database.GetDB()
    var foodItems []models.FoodItem

    if err := db.Where("restaurant_id = ?", restaurentID).Find(&foodItems).Error; err != nil{
        return nil, err
    }

    return &foodItems, nil
}

func AddFoodToCart(foodID any, userID any)(error){
    db = database.GetDB()
    var foodItem models.FoodItem
    if err := db.First(&foodItem, foodID).Error; err != nil {
        return err
    }

    var cartItem models.CartItem
    if err := db.Where("user_id = ? AND food_item_id = ? AND is_checked_out = ?", userID, foodItem.ID, false).First(&cartItem).Error; err != nil {
        newItem := models.CartItem{
            UserID:     userID.(uint),
            FoodItemID: foodItem.ID,
            Quantity:   1,
            Price:     foodItem.Price,
        }
        if err := db.Create(&newItem).Error; err != nil {
            return err
        }
    } else {
        cartItem.Quantity++
        cartItem.Price += foodItem.Price
        if err := db.Save(&cartItem).Error; err != nil {
            return err
        }
    }
    return nil
}

func RemoveFoodFromCart(foodID any, userID any)(error){
    db = database.GetDB()
    var foodItem models.FoodItem
    if err := db.First(&foodItem, foodID).Error; err != nil {
        return err
    }

    var cartItem models.CartItem
    if err := db.Where("user_id = ? AND food_item_id = ? AND is_checked_out = ?", userID, foodItem.ID, false).First(&cartItem).Error; err != nil {
        return err
    }

    cartItem.Quantity--
    cartItem.Price -= foodItem.Price
    if cartItem.Quantity == 0 {
        if err := db.Delete(&cartItem).Error; err != nil {
            return err
        }
    } else {
        if err := db.Save(&cartItem).Error; err != nil {
            return err
        }
    }

    return nil
}