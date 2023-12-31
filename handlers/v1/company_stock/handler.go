package v1companystock

import (
	"capital-challenge-server/dbHelper"
	"capital-challenge-server/models"
	"capital-challenge-server/polygon"
	"database/sql"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/u2takey/go-utils/uuid"
)

const (
	Buy  = "buy"
	Sell = "sell"
	DefaultPageSize = "10"
	DefaultPage = "1"
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

// BuyCompanyStock godoc
// @Summary      BuyCompanyStock
// @Description  buy the selected amount of company stock
// @Tags         company_stock
// @Param        user_id   path      string  true  "user_id"
// @Param        request   body      BuyCompanyStockRequest  true  "request"
// @Success      200 {object} CompanyStockTransactionResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /company-stocks/buy/{user_id} [post]
func BuyCompanyStock(c *gin.Context) {
	userID := c.Param("user_id")
	user, err := dbHelper.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req BuyCompanyStockRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = validateBuyCompanyStocksRequest(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userBalance, companyStock, err := getUserBalanceAndCompanyStock(userID, req.CompanyStockID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if there is enough funds to buy stock
	if userBalance.CurrentBalance < (companyStock.VolumeWeightedAveragePrice * float32(req.Quantity)) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "insufficient funds"})
		return
	}

	updatedBalance := calculateUserBalance(userBalance.CurrentBalance, req.Quantity, companyStock.VolumeWeightedAveragePrice, Buy)
	
	err = dbHelper.UpdateUserBalance(user.ID, updatedBalance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	var userTransaction CompanyStockTransactionResponse
	userTransaction.UserTransaction.ID = uuid.NewUUID()
	userTransaction.UserTransaction.UserID = user.ID
	userTransaction.UserTransaction.CompanyStockID = companyStock.ID
	userTransaction.UserTransaction.Quantity = req.Quantity
	userTransaction.UserTransaction.BuyPrice = companyStock.VolumeWeightedAveragePrice
	userTransaction.CurrentBalance = updatedBalance

	err = dbHelper.InsertUserTransaction(userTransaction.UserTransaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, userTransaction)
		return

	} else if err != sql.ErrNoRows && err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	currentQuantity := calculateStockQuantity(userAssets.Quantity, req.Quantity, Buy)
	err = dbHelper.UpdateUserAssets(userAssets.ID, currentQuantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
// @Success      200 {object} CompanyStockTransactionResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /company-stocks/sell/{user_id} [post]
func SellCompanyStock(c *gin.Context) {
	userID := c.Param("user_id")
	user, err := dbHelper.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req SellCompanyStockRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = validateSellCompanyStocksRequest(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userBalance, companyStock, err := getUserBalanceAndCompanyStock(userID, req.CompanyStockID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userAssets, err := dbHelper.GetUserAssetsByCompanyStock(companyStock.ID, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if req.Quantity > userAssets.Quantity {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "quantity to sell is more than number of owned stock"})
		return
	}

	updatedBalance := calculateUserBalance(userBalance.CurrentBalance, req.Quantity, companyStock.VolumeWeightedAveragePrice, Sell)
	currentQuantity := calculateStockQuantity(userAssets.Quantity, req.Quantity, Sell)
	if currentQuantity == 0 {
		err := dbHelper.DeleteUserAssets(userAssets.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		err = dbHelper.UpdateUserAssets(userAssets.ID, currentQuantity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	var userTransaction CompanyStockTransactionResponse
	userTransaction.UserTransaction.ID = uuid.NewUUID()
	userTransaction.UserTransaction.UserID = user.ID
	userTransaction.UserTransaction.CompanyStockID = companyStock.ID
	userTransaction.UserTransaction.Quantity = req.Quantity
	userTransaction.UserTransaction.SellPrice = companyStock.VolumeWeightedAveragePrice
	userTransaction.CurrentBalance = updatedBalance
	err = dbHelper.InsertUserTransaction(userTransaction.UserTransaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = dbHelper.UpdateUserBalance(user.ID, updatedBalance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userTransaction)
}

func calculateUserBalance(currentBalance float32, quantity int, volumeWeightedAveragePrice float32, operation string)(float32){
	var updatedBalance float32
	if operation == Sell {
		updatedBalance = currentBalance + (float32(quantity) * volumeWeightedAveragePrice)
	} else {
		updatedBalance = currentBalance - (float32(quantity) * volumeWeightedAveragePrice)
	} 

	return updatedBalance
}

func calculateStockQuantity(currentQuantity int, requestedQuantity int, operation string)(int){
	if operation == Sell{
		return currentQuantity - requestedQuantity
	} else {
		return currentQuantity + requestedQuantity
	}
}

func getUserBalanceAndCompanyStock(userID string, companyStockID string)(models.UserBalance, models.CompanyStock, error){
	userBalance, err := dbHelper.GetUserBalanceByUserID(userID)
	if err != nil {
		return userBalance , models.CompanyStock{}, err
	}

	companyStock, err := dbHelper.GetCompanyStockByID(companyStockID)
	if err != nil {
		return userBalance, companyStock, err
	}

	return userBalance , companyStock, nil
}

// GetDailyCompanyStockList godoc
// @Summary      GetDailyCompanyStockList
// @Description  get daily company stock list
// @Tags         company_stock
// @Param        page   query      string  false  "page" Format(string)
// @Param        page_size  query       string  false  "page_size" Format(string)
// @Success      200 {object} []models.CompanyStockList
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /company-stocks/list [get]
func GetDailyCompanyStockList(c *gin.Context){
	var params CompanyStockListQueryParams
	params.Page = c.Query("page")
	params.PageSize = c.Query("page_size")
	pageInt, pageSizeInt, err := castStringParamsToInt(params)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	list, err := dbHelper.GetCompanyStockList(*pageInt, *pageSizeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK,list)
}


// GetHistoricValueOfCompanyStock godoc
// @Summary      GetHistoricValueOfCompanyStock
// @Description  get history value of company stock
// @Tags         company_stock
// @Param        ticker   path      string  true  "ticker"
// @Success      200 {object} []models.CompanyStock
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /company-stocks/{ticker} [get]
func GetHistoricValueOfCompanyStock(c *gin.Context){
	var ticker string
	ticker = c.Param("ticker")
	list, err := dbHelper.GetHistoricalValueOfCompanyStock(ticker)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if list == nil {
		c.JSON(http.StatusOK, gin.H{"message": "No data for required ticker"})
		return
	} else {
		c.JSON(http.StatusOK,list)
	}
}

func castStringParamsToInt(params CompanyStockListQueryParams)(*int, *int, error){

	if params.Page == ""{
		params.Page = DefaultPage
	}

	if params.PageSize == ""{
		params.PageSize = DefaultPageSize
	}

	page, err := strconv.Atoi(params.Page)
	if err != nil {
		return nil, nil, err
	}
	pageSize, err := strconv.Atoi(params.PageSize)
	if err != nil {
		return nil , nil, err
	}
	return &page, &pageSize, nil
}