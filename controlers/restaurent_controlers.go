package controlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rishabhvegrow/foodx-go-server/models"
	"github.com/rishabhvegrow/foodx-go-server/repositories"
)

// Restaurents
func GetRestaurents(c *gin.Context) {
    restaurents, err := repositories.GetRestaurents()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Restaurents"})
        return
    }
    c.JSON(http.StatusOK, restaurents)
}

func CreateRestaurent(c *gin.Context) {
    var restaurent models.Restaurant
    if err := c.ShouldBindJSON(&restaurent); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    createdRestaurent, err := repositories.CraeteRestaurent(restaurent)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create restaurent"})
        return
    }
    c.JSON(http.StatusCreated, createdRestaurent)
}

func DeleteRestaurent(c *gin.Context) {
    id := c.Param("id")
    err := repositories.DeleteRestaurent(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Restaurent not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Restaurent deleted"})
}

func GetRestaurent(c *gin.Context) {
    id := c.Param("id")
    restaurent, err := repositories.GetRestaurent(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Restaurent not found"})
        return
    }
    c.JSON(http.StatusOK, restaurent)
}

