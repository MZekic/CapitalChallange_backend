package v1userassets

import (
	"capital-challenge-server/dbHelper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserAssets godoc
// @Summary      GetUserAssets
// @Description  get user assets for current game
// @Tags         user_assets
// @Param        user_id   path      string  true  "user_id"
// @Success      200 {object} []models.UserAssets
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /user-assets/{user_id} [get]
func GetUserAssets(c *gin.Context) {
	userID := c.Param("user_id")
	user, err := dbHelper.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userAssets, err := dbHelper.GetUserAssets(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if userAssets == nil {
		c.JSON(http.StatusOK, gin.H{"message": "no data"})
	} else {
		c.JSON(http.StatusOK, userAssets)
	}
}

// GetUserAssetsValue godoc
// @Summary      GetUserAssetsValue
// @Description  get user assets values
// @Tags         user_assets
// @Param        user_id   path      string  true  "user_id"
// @Success      200 {object} UserAssetsValues
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /user-assets/value/{user_id} [get]
func GetUserAssetsValue(c *gin.Context) {
	userID := c.Param("user_id")
	user, err := dbHelper.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userAssets, err := dbHelper.GetUserAssetsQuantityByTicker(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var values []UserAssetsValues
	for _, data := range userAssets {
		var unitValue UserAssetsValues
		asset, err := dbHelper.GetCurrentCompanyStockByTicker(data.Ticker)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		unitValue.Ticker = data.Ticker
		unitValue.ValuePerUnit = asset.VolumeWeightedAveragePrice
		unitValue.Value = asset.VolumeWeightedAveragePrice * float32(data.Quantity)
		unitValue.Quantity = data.Quantity

		values = append(values, unitValue)

	}

	if values == nil {
		c.JSON(http.StatusOK, gin.H{"message": "no data"})
	} else {
		c.JSON(http.StatusOK, values)
	}

}

// GetUserAssetsProfits godoc
// @Summary      GetUserAssetsProfits
// @Description  get user assets profits
// @Tags         user_assets
// @Param        user_id   path      string  true  "user_id"
// @Success      200 {object} UserAssetsProfits
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /user-assets/profits/{user_id} [get]
func GetUserProfits(c *gin.Context) {
	userID := c.Param("user_id")
	user, err := dbHelper.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	userAssets, err := dbHelper.GetUserAssets(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userProfits UserAssetsProfits

	for _, data := range userAssets {
		var userAssetsValues UserAssetsStocksInfo
		companyStock, err := dbHelper.GetCompanyStockByID(data.CompanyStockID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		currentStockValue, err := dbHelper.GetCurrentCompanyStockByTicker(data.Ticker)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userAssetsValues.CompanyStock = companyStock
		userAssetsValues.Quantity = data.Quantity
		userAssetsValues.BuyPrice = data.PricePerUnit
		userAssetsValues.CurrentPrice = currentStockValue.VolumeWeightedAveragePrice
		userAssetsValues.TotalBuyPrice = float32(data.Quantity) * data.PricePerUnit
		userAssetsValues.TotalCurrentPrice = float32(data.Quantity) * currentStockValue.VolumeWeightedAveragePrice
		userAssetsValues.ProfitPerUnit = currentStockValue.VolumeWeightedAveragePrice - data.PricePerUnit
		userAssetsValues.TotalProfit = float32(data.Quantity) * userAssetsValues.ProfitPerUnit
		userAssetsValues.ProfitMargin = calculateProfitMargin(data.PricePerUnit, currentStockValue.VolumeWeightedAveragePrice)
		userProfits.UserAssets = append(userProfits.UserAssets, userAssetsValues)
		userProfits.TotalSpent += userAssetsValues.TotalBuyPrice
		userProfits.TotalCurrentValue += userAssetsValues.TotalCurrentPrice
	}
	userProfits.TotalProfitMargin = calculateProfitMargin(userProfits.TotalSpent,userProfits.TotalCurrentValue)
	userProfits.TotalProfit = userProfits.TotalCurrentValue - userProfits.TotalSpent

	if userProfits.UserAssets == nil {
		c.JSON(http.StatusOK, gin.H{"message": "no data"})
	} else {
		c.JSON(http.StatusOK, userProfits)
	}
}

func calculateProfitMargin(originalPrice float32, currentPrice float32)(string){
	return fmt.Sprintf("%f%%", ((currentPrice - originalPrice) / originalPrice * 100))
}
