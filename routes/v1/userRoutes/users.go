package v1UserRoutes

import (
	v1Users "capital-challenge-server/handlers/v1/users"

	"github.com/gin-gonic/gin"
)

func AddUserRoutes(rg *gin.RouterGroup) {
	rg.POST("/register", v1Users.Registration)
	rg.POST("/", v1Users.Login)
	rg.GET("/:user_id", v1Users.GetUser)
}