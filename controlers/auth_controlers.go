package controlers

import (
	"net/http"
    "gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/rishabhvegrow/foodx-go-server/database"
	"github.com/rishabhvegrow/foodx-go-server/models"
)

var db *gorm.DB


// Login
func Login(c *gin.Context){
    db = database.GetDB()
    var requestData map[string]interface{}
    if err := c.BindJSON(&requestData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    email, emailExists := requestData["email"].(string)
    password, passwordExists := requestData["password"].(string)

    if !emailExists || !passwordExists {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
        return
    }

    var user models.User
    if err := db.Where("email = ?", email).First(&user).Error; err != nil || user.Password != password {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, err := generateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token, "email": user.Email, "name": user.Name})
}