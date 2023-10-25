package controlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rishabhvegrow/foodx-go-server/models"
	"github.com/rishabhvegrow/foodx-go-server/repositories"
)


func GetUsers(c *gin.Context) {
    users, err := repositories.GetUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
        return
    }
    c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
    id := c.Param("id")
    user, err := repositories.GetUser(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    createdUser, err := repositories.CraeteUser(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }
    c.JSON(http.StatusCreated, createdUser)
}

func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    
    err := repositories.DeleteUser(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}