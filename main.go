package main

import (
	"log"
	"os"

	database "capital-challenge-server/database"
	polygon "capital-challenge-server/polygon"
	routes "capital-challenge-server/routes"

	_ "capital-challenge-server/docs"

	"github.com/joho/godotenv"
	"github.com/robfig/cron"
)

// @title           Capital-Challenge API
// @version         1.0
// @description     Capital-Challenge backend API.

// @host      localhost:8080
// @BasePath  /v1

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						x-api-key
//	@description				x-api-key

// @Security ApiKeyAuth
func main() {
	_, dev := os.LookupEnv("ENVIROMENT")
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

	cronJob := cron.New()
	cronJob.AddFunc("* * * * *", justPrint)

	cronJob.Start()

	routes.Run()

}

func justPrint() {
	log.Println("JOB JOB JOB")
}
