package routes

import (
	_ "capital-challenge-server/docs"
	"capital-challenge-server/middlewares"
	v1CompaniesRoutes "capital-challenge-server/routes/v1/companiesRoutes"
	v1CompanyStocksRoutes "capital-challenge-server/routes/v1/company_stock"
	v1UserAssetsRoutes "capital-challenge-server/routes/v1/userAssets"
	v1UsersRoutes "capital-challenge-server/routes/v1/userRoutes"
	v1UserTransactionsRoutes "capital-challenge-server/routes/v1/userTransactionsRoutes"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var router = gin.Default()

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x-api-key")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Run() {
	router.Use(CORS())
	getRoutes()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run()
}

func getRoutes() {
	v1Companies := router.Group("/v1/companies", middlewares.Auth)
	v1CompaniesRoutes.AddCompaniesRoutes(v1Companies)

	v1CompanyStocks := router.Group("/v1/company-stocks", middlewares.Auth)
	v1CompanyStocksRoutes.AddCompanyStocksRoutes(v1CompanyStocks)

	v1Users := router.Group("/v1/users", middlewares.Auth)
	v1UsersRoutes.AddUserRoutes(v1Users)

	v1UserAssets := router.Group("/v1/user-assets", middlewares.Auth)
	v1UserAssetsRoutes.AddUserAssetsRoutes(v1UserAssets)

	v1UserTransactions := router.Group("/v1/user-transactions", middlewares.Auth)
	v1UserTransactionsRoutes.AddUserTransactions(v1UserTransactions)
}
