package dbHelper

import (
	database "capital-challenge-server/database"
	"capital-challenge-server/models"
)


//check buy or sell in query to assign null value
func InsertUserTransaction(userTransaction models.UserTransactions) error {
	sqlQuery := "INSERT INTO user_transactions (id, user_id, company_stock_id, quantity, buy_price, sell_price) VALUES (:id, :user_id, :company_stock_id, :quantity, :buy_price, :sell_price)"
	_, err := database.DB.NamedExec(sqlQuery, userTransaction)
	if err != nil {
		return err
	}
	return nil
}

func GetUserTransactions(userID string) ([]models.UserTransactionUnit, error) {
	sqlQuery := "SELECT t.quantity, t.buy_price, t.sell_price, t.company_stock_id, t.created_at, s.ticker FROM user_transactions t JOIN company_stock s ON s.id = t.company_stock_id WHERE t.user_id = $1 ORDER BY t.created_at"

	var userTransactions []models.UserTransactionUnit

	err := database.DB.Select(&userTransactions, sqlQuery, userID)

	if err != nil {
		return userTransactions, err
	}

	return userTransactions, nil
}


func GetUserTransactionsByTicker(userID string, ticker string)([]models.UserTransactionUnit, error){
	sqlQuery := "SELECT t.quantity, t.buy_price, t.sell_price, t.company_stock_id, t.created_at, s.ticker FROM user_transactions t JOIN company_stock s ON s.id = t.company_stock_id WHERE t.user_id = $1 AND s.ticker = $2 ORDER BY t.created_at"

	var userTransactions []models.UserTransactionUnit

	err := database.DB.Select(&userTransactions, sqlQuery, userID, ticker)

	if err != nil {
		return userTransactions, err
	}

	return userTransactions, nil
}



