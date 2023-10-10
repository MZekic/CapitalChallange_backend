package dbHelper

import (
	"capital-challenge-server/database"
	"capital-challenge-server/models"
)

func BulkInsertCompanyStock(companyStocks []models.CompanyStock) error {
	if len(companyStocks) > 0 {
		//todo test this shit
		companyStocksLen := len(companyStocks)
		companyStocks1 := companyStocks[:companyStocksLen/2]
		companyStocks2 := companyStocks[(companyStocksLen/2):]

		sqlQuery := `INSERT INTO company_stock (id, ticker, close_price, highest_price, lowest_price, number_of_transactions, open_price, otc, trading_volume, volume_weighted_average_price) VALUES 
	(:id, :ticker, :close_price, :highest_price, :lowest_price, :number_of_transactions, :open_price, :otc, :trading_volume, :volume_weighted_average_price)`
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
	sqlQuery := `INSERT INTO user_transactions (id, user_id, company_stock_id, buy_or_sell, quantity, game_number) 
	VALUES (:id, :user_id, :company_stock_id, :buy_or_sell, :quantity, :game_number)`

	_, err := database.DB.NamedExec(sqlQuery, buyRequest)
	if err != nil {
		return err
	}
	return nil
}

func GetCompanyStockInfoByID(companyStockID string) (models.CompanyStock, error) {

	var companyStock models.CompanyStock

	sqlQuery := "SELECT * FROM company_stock WHERE id = $1"
	err := database.DB.Get(&companyStock, sqlQuery, companyStockID)
	if err != nil {
		return companyStock, err
	}

	return companyStock, nil
}
