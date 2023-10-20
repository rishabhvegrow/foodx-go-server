package controlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/rishabhvegrow/foodx-go-server/database"
	"github.com/rishabhvegrow/foodx-go-server/models"
)

// Food Item
func CreateFoodItem(c *gin.Context){
    db = database.GetDB()
    var foodItem models.FoodItem
    if err := c.ShouldBindJSON(&foodItem); err != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db := db.Create(&foodItem)

    if db.Error != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create restaurent"})
        return
    }
    c.JSON(http.StatusCreated, foodItem)
}

func UpdateFoodItem(c *gin.Context){
    db = database.GetDB()
    id := c.Param("id")
    var foodItem models.FoodItem
    db := db.First(&foodItem, id)
    if db.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find record"})
        return
    }

    if err := c.ShouldBindJSON(&foodItem); err != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db.Save(&foodItem)

    if db.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update food item"})
        return
    }
    
    c.JSON(http.StatusOK, foodItem)
}

func DeleteFoodItem(c *gin.Context) {
    db = database.GetDB()
    id := c.Param("id")
    var foodItem models.FoodItem
    db := db.Delete(&foodItem, id)
    if db.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "foodItem not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "foodItem deleted"})
}

func GetFoodItemOfRestaurent(c *gin.Context){
    db = database.GetDB()
    restid := c.Param("restid")
    var foodItems []models.FoodItem

    if db.Where("restaurant_id = ?", restid).Find(&foodItems).Error != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find Records"})
    }
    c.JSON(http.StatusOK, foodItems)
}

func AddFoodToCart(c *gin.Context){
    db = database.GetDB()
    foodID := c.Param("id")
    userID := c.MustGet("user_id").(uint)

    var foodItem models.FoodItem
    if err := db.First(&foodItem, foodID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Food item not found"})
        return
    }

    var cartItem models.CartItem
    if err := db.Where("user_id = ? AND food_item_id = ? AND is_checked_out = ?", userID, foodItem.ID, false).First(&cartItem).Error; err != nil {
        // Create a cart item
        newItem := models.CartItem{
            UserID:     userID,
            FoodItemID: foodItem.ID,
            Quantity:   1,
            Price:     foodItem.Price,
        }
        if err := db.Create(&newItem).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to cart"})
            return
        }
    } else {
        cartItem.Quantity++
        cartItem.Price += foodItem.Price
        if err := db.Save(&cartItem).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart item"})
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{"message": "Item added to cart successfully"})
}

func RemoveFoodFromCart(c *gin.Context){
    db = database.GetDB()
    foodID := c.Param("id")
    userID := c.MustGet("user_id").(uint)

    var foodItem models.FoodItem
    if err := db.First(&foodItem, foodID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Food item not found"})
        return
    }

    var cartItem models.CartItem
    if err := db.Where("user_id = ? AND food_item_id = ? AND is_checked_out = ?", userID, foodItem.ID, false).First(&cartItem).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Item does not exist in cart"})
        return
    }

    cartItem.Quantity--
    cartItem.Price -= foodItem.Price
    if cartItem.Quantity == 0 {
        if err := db.Delete(&cartItem).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item from cart"})
            return
        }
    } else {
        if err := db.Save(&cartItem).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart item"})
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart successfully"})
}