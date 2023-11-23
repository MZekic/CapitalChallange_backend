package v1CompanyStockRoutes

import (
	v1CompanyStock "capital-challenge-server/handlers/v1/company_stock"

	"github.com/gin-gonic/gin"
)

func AddCompanyStocksRoutes(rg *gin.RouterGroup) {
	rg.GET("/", v1CompanyStock.GetDailyCompanyStock)
	rg.GET("/list", v1CompanyStock.GetDailyCompanyStockList)
	rg.GET("/:ticker", v1CompanyStock.GetHistoricValueOfCompanyStock)
	rg.POST("/buy/:user_id", v1CompanyStock.BuyCompanyStock)
	rg.POST("/sell/:user_id", v1CompanyStock.SellCompanyStock)
}
