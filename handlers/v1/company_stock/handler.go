package v1companystock

import (
	"capital-challenge-server/dbHelper"
	"capital-challenge-server/polygon"
	utils "capital-challenge-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetDailyCompanyStock godoc
// @Summary      Get Daily Company Stock
// @Description  get the daily information about company stocks
// @Tags         company_stock
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /company-stock [get]
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
