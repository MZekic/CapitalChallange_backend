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
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = dbHelper.BatchInsertCompanyStock(res)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

// BuyOrSellCompanyStock godoc
// @Summary      BuyOrSellCompanyStock
// @Description  buy or sell the selected amount of company stock
// @Tags         company_stock
// @Param        user_id   path      string  true  "user_id"
// @Param        request   body      BuyOrSellCompanyStocksRequest  true  "request"
// @Success      200 {object} models.UserTransactions
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /company-stocks/{user_id} [post]
func BuyOrSellCompanyStocks(c *gin.Context) {

	userID := c.Param("user_id")
	user, err := dbHelper.GetUserByID(userID)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
		utils.Log(c, http.StatusBadRequest)
		return
	}
	var req BuyOrSellCompanyStocksRequest
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		utils.Log(c, http.StatusBadRequest)
		return
	}
	err = validateBuyOrSellCompanyStocksRequest(req)
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
	if req.BuyOrSell == Buy {
		if userBalance.CurrentBalance < (companyStock.VolumeWeightedAveragePrice * float32(req.Quantity)) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "insufficient funds"})
			utils.Log(c, http.StatusBadRequest)
			return
		}
	}
	var userTransaction models.UserTransactions
	userTransaction.ID = uuid.NewUUID()
	userTransaction.UserID = user.ID
	userTransaction.CompanyStockID = companyStock.ID
	userTransaction.Quantity = req.Quantity

	err = dbHelper.InsertUserTransactionOnBuy(userTransaction)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		utils.Log(c, http.StatusInternalServerError)
		return
	}
	var updatedBalance float32
	if req.BuyOrSell == Buy {
		updatedBalance = userBalance.CurrentBalance - (float32(req.Quantity) * companyStock.VolumeWeightedAveragePrice)
		err = dbHelper.UpdateUserBalanceOnUserTransaction(user.ID, updatedBalance)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Println(err)
			utils.Log(c, http.StatusInternalServerError)
			return
		}
	} else if req.BuyOrSell == Sell {
		updatedBalance = userBalance.CurrentBalance + (companyStock.VolumeWeightedAveragePrice * float32(req.Quantity))
		err = dbHelper.UpdateUserBalanceOnUserTransaction(user.ID, updatedBalance)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Println(err)
			utils.Log(c, http.StatusInternalServerError)
			return
		}
	}

	userAssets, err := dbHelper.GetUserAssetsByTicker(companyStock.Ticker, user.ID)

	if err == sql.ErrNoRows {
		var userAssets models.UserAssets
		userAssets.ID = uuid.NewUUID()
		userAssets.Ticker = companyStock.Ticker
		userAssets.Quantity = req.Quantity
		userAssets.UserID = user.ID
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
	var currentQuantity int
	if req.BuyOrSell == Buy {
		currentQuantity = userAssets.Quantity + req.Quantity
	} else if req.BuyOrSell == Sell {
		currentQuantity = userAssets.Quantity - req.Quantity
	}
	err = dbHelper.UpdateUserAssetsOnUserTransaction(userAssets.ID, currentQuantity)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		utils.Log(c, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, userTransaction)
	utils.Log(c, http.StatusOK)
}

// BuyCompanyStock godoc
// @Summary      BuyCompanyStock
// @Description  buy the selected amount of company stock
// @Tags         company_stock
// @Param        user_id   path      string  true  "user_id"
// @Param        request   body      BuyCompanyStockRequest  true  "request"
// @Success      200 {object} models.UserTransactions
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /company-stocks/buy/{user_id} [post]
func BuyCompanyStock(c *gin.Context) {
	userID := c.Param("user_id")
	user, err := dbHelper.GetUserByID(userID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var req BuyCompanyStockRequest
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = validateBuyCompanyStocksRequest(req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userBalance, err := dbHelper.GetUserBalanceByUserID(userID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	companyStock, err := dbHelper.GetCompanyStockInfoByID(req.CompanyStockID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if userBalance.CurrentBalance < (companyStock.VolumeWeightedAveragePrice * float32(req.Quantity)) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "insufficient funds"})
		return
	}

	var userTransaction models.UserTransactions
	userTransaction.ID = uuid.NewUUID()
	userTransaction.UserID = user.ID
	userTransaction.CompanyStockID = companyStock.ID
	userTransaction.Quantity = req.Quantity
	userTransaction.BuyPrice = companyStock.VolumeWeightedAveragePrice

	err = dbHelper.InsertUserTransactionOnBuy(userTransaction)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	updatedBalance := userBalance.CurrentBalance - (float32(req.Quantity) * companyStock.VolumeWeightedAveragePrice)
	err = dbHelper.UpdateUserBalanceOnUserTransaction(user.ID, updatedBalance)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	userAssets, err := dbHelper.GetUserAssetsByCompanyStock(companyStock.ID, user.ID)
	if err == sql.ErrNoRows {
		var userAssets models.UserAssets
		userAssets.ID = uuid.NewUUID()
		userAssets.Ticker = companyStock.Ticker
		userAssets.Quantity = req.Quantity
		userAssets.UserID = user.ID
		userAssets.CompanyStockID = req.CompanyStockID
		userAssets.PricePerUnit = companyStock.VolumeWeightedAveragePrice
		err = dbHelper.InsertUserAsset(userAssets)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, userTransaction)
		return

	} else if err != sql.ErrNoRows && err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	currentQuantity := userAssets.Quantity + req.Quantity
	err = dbHelper.UpdateUserAssetsOnUserTransaction(userAssets.ID, currentQuantity)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, userTransaction)
}

// SellCompanyStock godoc
// @Summary      SellCompanyStock
// @Description  sell the selected amount of company stock
// @Tags         company_stock
// @Param        user_id   path      string  true  "user_id"
// @Param        request   body      SellCompanyStockRequest  true  "request"
// @Success      200 {object} models.UserTransactions
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /company-stocks/sell/{user_id} [post]
func SellCompanyStock(c *gin.Context) {
	userID := c.Param("user_id")
	user, err := dbHelper.GetUserByID(userID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var req SellCompanyStockRequest
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = validateSellCompanyStocksRequest(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userBalance, err := dbHelper.GetUserBalanceByUserID(userID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	companyStock, err := dbHelper.GetCompanyStockInfoByID(req.CompanyStockID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userAssets, err := dbHelper.GetUserAssetsByCompanyStock(companyStock.ID, user.ID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	//AAAAAAAAAAAAAAAAAAAAAAAAAAAA
	if req.Quantity > userAssets.Quantity {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "quantity to sell is more than number of owned stock"})
		return
	}
	//AAAAAAAAAAAAAAAAAAAAAAA
	updatedBalance := userBalance.CurrentBalance + (float32(req.Quantity) * companyStock.VolumeWeightedAveragePrice)
	currentQuantity := userAssets.Quantity - req.Quantity
	if currentQuantity == 0 {
		err := dbHelper.DeleteUserAssets(userAssets.ID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	} else {
		err = dbHelper.UpdateUserAssets(userAssets.ID, currentQuantity)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	var userTransaction models.UserTransactions
	userTransaction.ID = uuid.NewUUID()
	userTransaction.UserID = user.ID
	userTransaction.CompanyStockID = companyStock.ID
	userTransaction.Quantity = req.Quantity
	userTransaction.SellPrice = companyStock.VolumeWeightedAveragePrice

	err = dbHelper.InsertUserTransactionOnSell(userTransaction)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = dbHelper.UpdateUserBalanceOnUserTransaction(user.ID, updatedBalance)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, userTransaction)
}
