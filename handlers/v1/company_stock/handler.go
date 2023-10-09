package v1companystock

import (
	"capital-challenge-server/dbHelper"
	"capital-challenge-server/models"
	"capital-challenge-server/polygon"
	utils "capital-challenge-server/utils"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/u2takey/go-utils/uuid"
)

const (
	Buy  = "buy"
	Sell = "sell"
)

// GetDailyCompanyStock godoc
// @Summary      Get Daily Company Stock
// @Description  get the daily information about company stocks
// @Tags         company_stock
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /company-stocks [get]
func GetDailyCompanyStock(c *gin.Context) {

	res, err := polygon.GetDailyCompanyStockInfo()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		utils.Log(c, http.StatusInternalServerError)
		return
	}

	err = dbHelper.BulkInsertCompanyStock(res)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		utils.Log(c, http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
	utils.Log(c, http.StatusOK)
}

// BuyCompanyStock godoc
// @Summary      BuyCompanyStock
// @Description  buy the selected amount of company stock
// @Tags         company_stock
// @Param        user_id   path      string  true  "user_id"
// @Param        request   body      BuyCompanyStocksRequest  true  "request"
// @Success      200 {object} models.UserTransactions
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /company-stocks/{user_id} [post]
func BuyCompanyStocks(c *gin.Context) {

	userID := c.Param("user_id")
	user, err := dbHelper.GetUserByID(userID)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
		utils.Log(c, http.StatusBadRequest)
		return
	}
	var req BuyCompanyStocksRequest
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		utils.Log(c, http.StatusBadRequest)
		return
	}
	err = validateBuyCompanyStocksRequest(req)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		utils.Log(c, http.StatusBadRequest)
		return
	}
	userBalance, err := dbHelper.GetUserBalanceByUserID(userID)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		utils.Log(c, http.StatusBadRequest)
		return
	}
	companyStock, err := dbHelper.GetCompanyStockInfoByID(req.CompanyStockID)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		utils.Log(c, http.StatusBadRequest)
		return
	}
	if userBalance.CurrentBalance < (companyStock.VolumeWeightedAveragePrice * float32(req.Quantity)) {

		c.JSON(http.StatusBadRequest, gin.H{"message": "insufficient funds"})
		utils.Log(c, http.StatusBadRequest)
		return
	}
	var userTransaction models.UserTransactions
	userTransaction.ID = uuid.NewUUID()
	userTransaction.BuyOrSell = Buy
	userTransaction.UserID = user.ID
	userTransaction.CompanyStockID = companyStock.ID
	userTransaction.GameNumber = user.CurrentGameNumber
	userTransaction.Quantity = req.Quantity

	err = dbHelper.InsertUserTransaction(userTransaction)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		utils.Log(c, http.StatusInternalServerError)
		return
	}
	updatedBalance := userBalance.CurrentBalance - (float32(req.Quantity) * companyStock.VolumeWeightedAveragePrice)

	err = dbHelper.UpdateUserBalanceOnBuy(user.ID, updatedBalance)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		utils.Log(c, http.StatusInternalServerError)
		return
	}

	userAssets, err := dbHelper.GetUserAssetsByTicker(companyStock.Ticker, user.ID)

	if err == sql.ErrNoRows {
		var userAssets models.UserAssets
		userAssets.ID = uuid.NewUUID()
		userAssets.Ticker = companyStock.Ticker
		userAssets.Quantity = req.Quantity
		userAssets.UserID = user.ID
		userAssets.GameNumber = user.CurrentGameNumber
		err = dbHelper.InsertUserAsset(userAssets)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Println(err)
			utils.Log(c, http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, userTransaction)
		utils.Log(c, http.StatusOK)
		return 

	} else if err != sql.ErrNoRows && err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		utils.Log(c, http.StatusInternalServerError)
		return
	}

	currentQuantity := userAssets.Quantity + req.Quantity
	err = dbHelper.UpdateUserAssetsOnBuyByTicker(userAssets.ID, currentQuantity)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		utils.Log(c, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, userTransaction)
	utils.Log(c, http.StatusOK)
}
