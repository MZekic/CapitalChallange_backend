package v1CompaniesRoutes

import (
	v1Companies "capital-challenge-server/handlers/v1/companies"

	"github.com/gin-gonic/gin"
)

func AddCompaniesRoutes(rg *gin.RouterGroup) {
	rg.GET("/:ticker", v1Companies.GetCompanyInfo)
}
