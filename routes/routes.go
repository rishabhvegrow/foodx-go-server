package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rishabhvegrow/foodx-go-server/controlers"
)


func SetupRoutes(router *gin.Engine) {

    base := router.Group("/")
    base.GET("/", controlers.GetPing)
    base.POST("register/", controlers.CreateUser)
    base.POST("login/", controlers.Login)

    userGroup := router.Group("/users")
    userGroup.GET("/", controlers.GetUsers)
    userGroup.GET("/:id", controlers.GetUser)
    userGroup.DELETE("/:id", controlers.DeleteUser)

    restGroup := router.Group("/restaurents")
    // restGroup.Use(controlers.JWTAuthMiddleware())
    restGroup.GET("/", controlers.GetRestaurents)
    restGroup.GET("/:id", controlers.GetRestaurent)
    restGroup.POST("/", controlers.CreateRestaurent)
    restGroup.DELETE("/:id", controlers.DeleteRestaurent)

    foodGroup := router.Group("/food")
    foodGroup.Use(controlers.JWTAuthMiddleware())
    foodGroup.GET("/:restid", controlers.GetFoodItemOfRestaurent)
    foodGroup.POST("/", controlers.CreateFoodItem)
    foodGroup.PUT("/:id", controlers.UpdateFoodItem)
    foodGroup.DELETE("/:id", controlers.DeleteFoodItem)
    foodGroup.POST("/:id/add", controlers.AddFoodToCart)
    foodGroup.POST("/:id/remove", controlers.RemoveFoodFromCart)

    cartGroup := router.Group("/cart")
    cartGroup.Use(controlers.JWTAuthMiddleware())
    cartGroup.GET("/", controlers.GetCartDetails)
    cartGroup.GET("/orders", controlers.GetOrderedItems)
    cartGroup.POST("/checkout", controlers.CheckoutCart)

    transactionGroup := router.Group("/transactions")
    transactionGroup.Use(controlers.JWTAuthMiddleware())
    transactionGroup.GET("/", controlers.GetTransactions)
}