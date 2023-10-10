package routes

import (
	_ "capital-challenge-server/docs"
	v1CompaniesRoutes "capital-challenge-server/routes/v1/companiesRoutes"
	v1CompanyStocksRoutes "capital-challenge-server/routes/v1/company_stock"
	v1UsersRoutes "capital-challenge-server/routes/v1/userRoutes"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var router = gin.Default()

func Run() {

	getRoutes()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run()
}

func getRoutes() {
	v1Companies := router.Group("/v1/companies")
	v1CompaniesRoutes.AddCompaniesRoutes(v1Companies)

	v1CompanyStocks := router.Group("/v1/company-stocks")
	v1CompanyStocksRoutes.AddCompanyStocksRoutes(v1CompanyStocks)

	v1Users := router.Group("/v1/users")
	v1UsersRoutes.AddUserRoutes(v1Users)
}
