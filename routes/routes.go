package routes

import (
	v1CompaniesRoutes "capital-challenge-server/routes/v1/companiesRoutes"

	_ "capital-challenge-server/docs"

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
}
