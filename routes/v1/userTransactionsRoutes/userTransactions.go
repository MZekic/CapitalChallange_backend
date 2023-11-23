package v1UserTransactionsRoutes

import (
	v1UserTransactions "capital-challenge-server/handlers/v1/user_transactions"

	"github.com/gin-gonic/gin"
)

func AddUserTransactions(rg *gin.RouterGroup) {
	rg.GET("/:user_id", v1UserTransactions.GetUserTransactions)
	rg.GET("/:user_id/:ticker", v1UserTransactions.GetUserTransactionsByTicker)

}