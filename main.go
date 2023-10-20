package main

import (
    "log"
	"github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
	"github.com/rishabhvegrow/foodx-go-server/database"
    "github.com/rishabhvegrow/foodx-go-server/models"
    "github.com/rishabhvegrow/foodx-go-server/routes"
)


func main() {
    router := gin.Default()

    err := godotenv.Load()
        if err != nil {
        log.Fatal("Error loading .env file")
    }

    database.ConnectDB()

    // Migrating DB models
    db := database.GetDB()
    db.AutoMigrate(&models.User{}, &models.Restaurant{}, &models.FoodItem{}, &models.CartItem{}, &models.Transaction{})

    routes.SetupRoutes(router)

    router.Run(":8080")
}
