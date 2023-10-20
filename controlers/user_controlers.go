package controlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/rishabhvegrow/foodx-go-server/database"
	"github.com/rishabhvegrow/foodx-go-server/models"
)


func GetUsers(c *gin.Context) {
    var users []models.User
    db = database.GetDB()
    if db.Find(&users).Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
        return
    }
    c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
    db = database.GetDB()
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db := db.Create(&user)
    if db.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }
    c.JSON(http.StatusCreated, user)
}

func DeleteUser(c *gin.Context) {
    db = database.GetDB()
    id := c.Param("id")
    var user models.User
    db := db.Delete(&user, id)
    if db.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func GetUser(c *gin.Context) {
    db = database.GetDB()
    id := c.Param("id")
    var user models.User
    db := db.First(&user, id)
    if db.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}