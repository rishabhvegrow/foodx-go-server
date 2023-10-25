package controlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/rishabhvegrow/foodx-go-server/repositories"
)

func RemoveCartItem(c *gin.Context) {
    
    id := c.Param("id")

    cartItem, err := repositories.GetCartItemByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "CartItem not found"})
        return
    }

    if !cartItem.IsCheckedOut {
        err = repositories.RemoveCartItemByID(cartItem.ID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete CartItem"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "CartItem deleted successfully"})
    } else {
        c.JSON(http.StatusForbidden, gin.H{"error": "CartItem cannot be deleted because it is checked out"})
    }
}

func GetCartDetails(c *gin.Context) {
    userID := c.MustGet("user_id").(uint)

    cartItems, err := repositories.GetCartItemsOfAUser(userID, false)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No items added in the cart"})
        return
    }

    c.JSON(http.StatusOK, cartItems)
}

func CheckoutCart(c *gin.Context) {

    userID := c.MustGet("user_id").(uint)

    cartItems, err := repositories.GetCartItemsOfAUser(userID, false)
    if err != nil || len(*cartItems) == 0 {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No items added to the cart"})
        return
    }

    transaction, err := repositories.CheckoutCart(userID)
    if err != nil || len(*cartItems) == 0 {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to checkout, Please try again"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Transaction successful", "transaction": transaction})
}

func GetTransactions(c *gin.Context){

    userID := c.MustGet("user_id").(uint)

    transactions, err := repositories.GetTransactionsOfUser(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
        return
    }
    c.JSON(http.StatusOK, transactions)
}

func GetOrderedItems(c *gin.Context){
    userID := c.MustGet("user_id").(uint)
    cartItems, err := repositories.GetCartItemsOfAUser(userID, true)
    
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No items added in the Orders"})
        return
    }

    c.JSON(http.StatusOK, cartItems)
}