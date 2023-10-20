package controlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/rishabhvegrow/foodx-go-server/database"
	"github.com/rishabhvegrow/foodx-go-server/models"
)

// Restaurents
func GetRestaurents(c *gin.Context) {
    db = database.GetDB()
    var restaurents []models.Restaurant
    db := db.Find(&restaurents)
    if db.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Restaurents"})
        return
    }
    c.JSON(http.StatusOK, restaurents)
}

func CreateRestaurent(c *gin.Context) {
    db = database.GetDB()
    var restaurent models.Restaurant
    if err := c.ShouldBindJSON(&restaurent); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db := db.Create(&restaurent)
    if db.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create restaurent"})
        return
    }
    c.JSON(http.StatusCreated, restaurent)
}

func DeleteRestaurent(c *gin.Context) {
    db = database.GetDB()
    id := c.Param("id")
    var restaurent models.Restaurant
    db := db.Delete(&restaurent, id)
    if db.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Restaurent not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Restaurent deleted"})
}

func GetRestaurent(c *gin.Context) {
    db = database.GetDB()
    id := c.Param("id")
    var restaurent models.Restaurant
    db := db.First(&restaurent, id)
    if db.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Restaurent not found"})
        return
    }
    c.JSON(http.StatusOK, restaurent)
}

func GetPing(c *gin.Context){
    c.JSON(http.StatusOK, "PONG")
}