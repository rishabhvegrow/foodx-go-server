package controlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rishabhvegrow/foodx-go-server/models"
	"github.com/rishabhvegrow/foodx-go-server/repositories"
)

// Food Item
func CreateFoodItem(c *gin.Context){
    var foodItem models.FoodItem
    if err := c.ShouldBindJSON(&foodItem); err != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
	createdFoodItem, err := repositories.CreateFoodItem(foodItem)

    if err != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create restaurent"})
        return
    }
    c.JSON(http.StatusCreated, createdFoodItem)
}

func UpdateFoodItem(c *gin.Context){
    
    id := c.Param("id")
    var foodItem models.FoodItem

    if err := c.ShouldBindJSON(&foodItem); err != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedFoodItem, err := repositories.UpdateFoodItem(id, foodItem)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update food item"})
        return
    }
    
    c.JSON(http.StatusOK, updatedFoodItem)
}

func DeleteFoodItem(c *gin.Context) {
    id := c.Param("id")
    err := repositories.DeleteFoodItem(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "foodItem not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "foodItem deleted"})
}

func GetFoodItemOfRestaurent(c *gin.Context){
    restid := c.Param("restid")
    foodItems, err := repositories.GetMenuOfRestaurent(restid)

    if err != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find Records"})
    }
    c.JSON(http.StatusOK, foodItems)
}

func AddFoodToCart(c *gin.Context){
    foodID := c.Param("id")
    userID := c.MustGet("user_id").(uint)

	err := repositories.AddFoodToCart(foodID, userID)

    if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart item"})
		return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Item added to cart successfully"})
}

func RemoveFoodFromCart(c *gin.Context){
    foodID := c.Param("id")
    userID := c.MustGet("user_id").(uint)

	err := repositories.RemoveFoodFromCart(foodID, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart item"})
		return
	}

    c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart successfully"})
}