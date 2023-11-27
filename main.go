package main

import (
	"log"
	"os"

	"capital-challenge-server/cron"
	database "capital-challenge-server/database"
	polygon "capital-challenge-server/polygon"
	routes "capital-challenge-server/routes"

	_ "capital-challenge-server/docs"

	"github.com/joho/godotenv"
)

// @title           Capital-Challenge API
// @version         1.0
// @description     Capital-Challenge backend API.

// @host      capital-challenge-server-nameless-cloud-7582.fly.dev
// @BasePath  /v1

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						x-api-key
//	@description				x-api-key

// @Security ApiKeyAuth
func main() {
	_, dev := os.LookupEnv("DEV")
	if !dev {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Some error occured. Err: %s", err)
		}
	}

	err := database.DBConnection()
	if err != nil {
		log.Fatalf("Error with initial connection to DB: %s", err)
	}

	polygon.StartPolygonClient()

	cron.GetDailyStocks()

	routes.Run()

}
