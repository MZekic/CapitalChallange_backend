package dbHelper

import (
	database "capital-challenge-server/database"
	"capital-challenge-server/models"
	"log"
)

func GetCompanyInfoByTicker(ticker string) (*models.Companies, error) {
	var result models.Companies
	log.Println("ASD: ", ticker)
	sqlQuery := "SELECT * FROM companies WHERE ticker = $1"
	err := database.DB.Get(&result, sqlQuery, ticker)
	if err != nil {
		return &result, err
	}

	return &result, nil
}
