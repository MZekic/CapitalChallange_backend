package dbHelper

import (
	database "capital-challenge-server/database"
	"capital-challenge-server/models"
)

func GetCompanyInfoByTicker(ticker string) (*models.Companies, error) {
	var result models.Companies
	sqlQuery := "SELECT * FROM companies WHERE ticker = $1"
	err := database.DB.Get(&result, sqlQuery, ticker)
	if err != nil {
		return &result, err
	}

	return &result, nil
}

func InsertCompany(company models.Companies) error {
	sqlQuery := `INSERT INTO companies (id, name, ticker, market, primary_exchange, type, currency_name, locale, homepage_url, description, total_employees, logo_url) VALUES 
	(:id, :name, :ticker, :market, :primary_exchange, :type, :currency_name, :locale, :homepage_url, :description, :total_employees, :logo_url)`

	_, err := database.DB.NamedExec(sqlQuery, company)
	if err != nil {
		return err
	}
	return nil
}
