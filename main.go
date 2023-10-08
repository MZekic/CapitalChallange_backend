package main

import (
	"log"

	database "capital-challenge-server/database"
	polygon "capital-challenge-server/polygon"
	routes "capital-challenge-server/routes"

	_ "capital-challenge-server/docs"

	"github.com/joho/godotenv"
)

// @title           Capital-Challenge API
// @version         1.0
// @description     Capital-Challenge backend API.

// @host      localhost:8080
// @BasePath  /v1

// @securityDefinitions.basic  BasicAuth

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
