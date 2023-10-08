package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Log(c *gin.Context, statusCode int) {
	f, err := os.OpenFile("server-logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	logMsg := fmt.Sprintf("%s: METHOD: %s, ROUTE: %s, STATUS CODE: %d, USER: %s, BODY: %s \n", time.Now(), c.Request.Method, c.Request.URL.Path, statusCode, "null", c.Request.Body)
	f.Write([]byte(logMsg))

}
