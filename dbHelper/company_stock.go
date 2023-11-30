package dbHelper

import (
	"capital-challenge-server/database"
	"capital-challenge-server/models"
)

func BatchInsertCompanyStock(companyStocks []models.CompanyStock) error {
	if len(companyStocks) > 0 {
		//todo test this shit
		companyStocksLen := len(companyStocks)
		companyStocks1 := companyStocks[:companyStocksLen/2]
		companyStocks2 := companyStocks[(companyStocksLen / 2):]

		sqlQuery := `INSERT INTO company_stock (id, ticker, close_price, highest_price, lowest_price, number_of_transactions, open_price, otc, trading_volume, volume_weighted_average_price, date) VALUES 
	(:id, :ticker, :close_price, :highest_price, :lowest_price, :number_of_transactions, :open_price, :otc, :trading_volume, :volume_weighted_average_price, :date)`
		_, err := database.DB.NamedExec(sqlQuery, companyStocks1)
		if err != nil {
			return err
		}
		_, err = database.DB.NamedExec(sqlQuery, companyStocks2)
		if err != nil {
			return err
		}
	}
	return nil
}

func BuyCompanyStocks(buyRequest models.CompanyStock) error {
	sqlQuery := `INSERT INTO user_transactions (id, user_id, company_stock_id, buy_or_sell, quantity) 
	VALUES (:id, :user_id, :company_stock_id, :buy_or_sell, :quantity)`

	_, err := database.DB.NamedExec(sqlQuery, buyRequest)
	if err != nil {
		return err
	}
	return nil
}

func GetCompanyStockByID(companyStockID string) (models.CompanyStock, error) {
	var companyStock models.CompanyStock

	sqlQuery := "SELECT * FROM company_stock WHERE id = $1"
	err := database.DB.Get(&companyStock, sqlQuery, companyStockID)
	if err != nil {
		return companyStock, err
	}

	return companyStock, nil
}
func GetCurrentCompanyStockByTicker(ticker string) (models.CompanyStock, error) {
	var companyStock models.CompanyStock

	sqlQuery := "SELECT * FROM company_stock WHERE ticker = $1 ORDER BY created_at DESC LIMIT 1"
	err := database.DB.Get(&companyStock, sqlQuery, ticker)
	if err != nil {
		return companyStock, err
	}

	return companyStock, nil
}

/*
func GetCompanyStockList(page int, pageSize int) (*models.CompanyStockList, error) {
	var companyStockList []models.CompanyStock

	sqlQuery := "SELECT * FROM company_stock WHERE date < now() - interval '1 day' and date > now() - interval '2 day' LIMIT $1 OFFSET $2"
	err := database.DB.Select(&companyStockList, sqlQuery, pageSize, (page*pageSize)-pageSize)
	if err != nil {
		return nil, err
	}
	sqlQueryCount := "SELECT COUNT(*) FROM company_stock WHERE date < now() - interval '1 day' and date > now() - interval '2 day'"
	var results int
	err = database.DB.Get(&results, sqlQueryCount)
	if err != nil {
		return nil, err
	}
	var companyStockListWithResultCount models.CompanyStockList
	companyStockListWithResultCount.CompanyStocks = companyStockList
	companyStockListWithResultCount.NumberOfResults = results

	return &companyStockListWithResultCount, nil
}*/

func GetCompanyStockList(page int, pageSize int) (*models.CompanyStockList, error) {
	var companyStockList []models.CompanyStock

	sqlQuery := "SELECT * FROM company_stock ORDER BY date, id DESC LIMIT $1 OFFSET $2"
	err := database.DB.Select(&companyStockList, sqlQuery, pageSize, (page*pageSize)-pageSize)
	if err != nil {
		return nil, err
	}
	/*
		sqlQueryCount := "SELECT COUNT(*) FROM company_stock WHERE date < now() - interval '1 day' and date > now() - interval '2 day'"
		var results int
		err = database.DB.Get(&results, sqlQueryCount)
		if err != nil {
			return nil, err
		}*/
	var companyStockListWithResultCount models.CompanyStockList
	companyStockListWithResultCount.CompanyStocks = companyStockList
	companyStockListWithResultCount.NumberOfResults = 90

	return &companyStockListWithResultCount, nil
}

func GetHistoricalValueOfCompanyStock(ticker string) ([]models.CompanyStock, error) {
	var companyStock []models.CompanyStock

	sqlQuery := "SELECT * FROM company_stock WHERE ticker = $1 ORDER BY created_at DESC"
	err := database.DB.Select(&companyStock, sqlQuery, ticker)
	if err != nil {
		return companyStock, err
	}
	return companyStock, nil
}
