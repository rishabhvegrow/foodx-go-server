package repositories

import (
	"github.com/rishabhvegrow/foodx-go-server/models"
)

func GetCartItemByID(id string)(*models.CartItem, error){

	var cartItem models.CartItem
	if err := db.Where("id = ?", id).First(&cartItem).Error; err != nil {
        return nil, err
    }

	return &cartItem, nil
}

func RemoveCartItemByID(id uint)(error) {
    
	if err := db.Where("id = ?", id).Delete(&models.CartItem{}).Error; err != nil {
        return err
    }
    return nil
}

func GetCartItemsOfAUser(userID uint, isChecked bool)(*[]models.CartItem, error){

	var cartItems []models.CartItem

    if err := db.
        Preload("FoodItem").
        Where("user_id = ? AND is_checked_out = ?", userID, isChecked).
        Find(&cartItems).Error; err != nil {
        return nil, err
    }

	return &cartItems, nil
}

func GetTransactionsOfUser(userID uint)(*models.Transaction, error){
	var transactions models.Transaction
	if err := db.Where("user_id = ?", userID).Find(&transactions).Error; err != nil{
		return nil, err
	}
	return &transactions, nil
}

func CheckoutCart(userID uint)(*models.Transaction, error){
	tx := db.Begin()
	if tx.Error != nil {
        return nil,tx.Error
    }
	cartItems, err := GetCartItemsOfAUser(userID, false)

	if err != nil || len(*cartItems) == 0 {
        tx.Rollback()
        return nil,err
	}

	transaction := models.Transaction{UserID: userID, Total: 0}
	if err := tx.Create(&transaction).Error; err != nil {
        tx.Rollback()
        return nil,err
    }

	var totalSum float32

    for _, item := range *cartItems {
        totalSum += item.Price
        item.IsCheckedOut = true
        item.TransactionID = transaction.ID
        if err := tx.Save(&item).Error; err != nil {
            tx.Rollback()
            return nil,err
        }
    }


    transaction.Total = float64(totalSum)
    if err := tx.Save(&transaction).Error; err != nil {
        tx.Rollback()
        return nil,err
    }

    tx.Commit()

	return &transaction,nil
}