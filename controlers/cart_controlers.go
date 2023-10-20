package controlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/rishabhvegrow/foodx-go-server/database"
	"github.com/rishabhvegrow/foodx-go-server/models"
)

//  Cart details
func GetCartDetails(c *gin.Context){
    db = database.GetDB()
    // Send all the cart items of the user
    userID := c.MustGet("user_id").(uint)

    var cartItems []models.CartItem
    if err := db.Where("user_id = ? AND is_checked_out = ?", userID, false).Find(&cartItems).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No items added in the cart"})
        return
    }

    c.JSON(http.StatusOK, cartItems)
}

func CheckoutCart(c *gin.Context) {
    db = database.GetDB()
    userID := c.MustGet("user_id").(uint)

    tx := db.Begin()
    if tx.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction start failed"})
        return
    }

    var cartItems []models.CartItem
    if err := tx.Where("user_id = ? AND is_checked_out = ?", userID, false).Find(&cartItems).Error; err != nil || len(cartItems) == 0 {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No items added to the cart"})
        return
    }

    transaction := models.Transaction{UserID: userID, Total: 0}
    if err := tx.Create(&transaction).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction creation failed"})
        return
    }

    var totalSum float32

    for _, item := range cartItems {
        totalSum += item.Price
        item.IsCheckedOut = true
        item.TransactionID = transaction.ID
        if err := tx.Save(&item).Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to checkout"})
            return
        }
    }


    transaction.Total = float64(totalSum)
    if err := tx.Save(&transaction).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction update failed"})
        return
    }

    tx.Commit()

    c.JSON(http.StatusOK, gin.H{"message": "Transaction successful", "transaction": transaction})
}

func GetTransactions(c *gin.Context){
    db = database.GetDB()
    userID := c.MustGet("user_id").(uint)

    var transactions []models.Transaction
    db := db.Where("user_id = ?", userID).Find(&transactions)
    if db.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
        return
    }
    c.JSON(http.StatusOK, transactions)
}

func GetOrderedItems(c *gin.Context){
    db = database.GetDB()
    userID := c.MustGet("user_id").(uint)

    var cartItems []models.CartItem
    db := db.Where("user_id = ? AND is_checked_out = ?", userID, true).Find(&cartItems)
    if db.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch ordered items"})
        return
    }
    c.JSON(http.StatusOK, cartItems)
}