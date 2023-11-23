package v1UserTransactions

import (
	"capital-challenge-server/dbHelper"
	"net/http"
	"github.com/gin-gonic/gin"
)


// GetUserTransactions godoc
// @Summary      GetUserTransactions
// @Description  get user transactions
// @Tags         user_transactions
// @Param        user_id   path      string  true  "user_id"
// @Success      200 {object} []models.UserTransactionUnit
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /user-transactions/{user_id} [get]
func GetUserTransactions(c *gin.Context){
	userID := c.Param("user_id")
	userTransactions, err := dbHelper.GetUserTransactions(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userTransactions == nil {
		c.JSON(http.StatusOK, gin.H{"message": "There are no transactions for this user yet"})
	} else {
		c.JSON(http.StatusOK, userTransactions)
	}
}

// GetUserTransactions godoc
// @Summary      GetUserTransactionsByTicker
// @Description  get user transactions by ticker
// @Tags         user_transactions
// @Param        user_id   path      string  true  "user_id"
// @Param        ticker   path      string  true  "ticker"
// @Success      200 {object} []models.UserTransactionUnit
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /user-transactions/{user_id}/{ticker} [get]
func GetUserTransactionsByTicker(c *gin.Context){
	userID := c.Param("user_id")
	ticker := c.Param("ticker")
	userTransactions, err := dbHelper.GetUserTransactionsByTicker(userID, ticker)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userTransactions == nil {
		c.JSON(http.StatusOK, gin.H{"message": "There are no transactions for this ticker yet"})
	} else {
		c.JSON(http.StatusOK, userTransactions)
	}
}
