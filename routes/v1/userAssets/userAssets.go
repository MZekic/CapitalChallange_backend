package v1userassetsroutes

import (
	v1userassets "capital-challenge-server/handlers/v1/user_assets"

	"github.com/gin-gonic/gin"
)

func AddUserAssetsRoutes(rg *gin.RouterGroup) {
	rg.GET("/:user_id", v1userassets.GetUserAssets)
	rg.GET("/value/:user_id", v1userassets.GetUserAssetsValue)
	rg.GET("/profits/:user_id", v1userassets.GetUserProfits)
}
