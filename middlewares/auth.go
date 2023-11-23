package middlewares

import (
	auth "capital-challenge-server/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func Auth(c *gin.Context) {
	if c.Request.Header.Get("x-api-key") != os.Getenv("X_API_KEY") {
	c.AbortWithStatus(http.StatusForbidden)
	}
}
