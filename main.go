package main

import (
	"log"

	database "capital-challenge-server/database"
	polygon "capital-challenge-server/polygon"
	routes "capital-challenge-server/routes"

	_ "capital-challenge-server/docs"

	"github.com/joho/godotenv"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	err = database.DBConnection()
	if err != nil {
		log.Fatalf("Error with initial connection to DB: %s", err)
	}

	polygon.StartPolygonClient()

	//r := gin.Default()
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//r.GET("/v1/companies/:ticker", v1Companies.GetCompanyInfo)
	//r.Run()
	routes.Run()
}
