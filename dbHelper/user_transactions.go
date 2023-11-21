package dbHelper

import (
	database "capital-challenge-server/database"
	"capital-challenge-server/models"
)


//check buy or sell in query to assign null value
func InsertUserTransaction(userTransaction models.UserTransactions) error {
	var sqlQuery string
	if userTransaction.BuyPrice == 0{
	sqlQuery = "INSERT INTO user_transactions (id, user_id, company_stock_id, quantity, sell_price) VALUES (:id, :user_id, :company_stock_id, :quantity, :sell_price)"
	} else {
		sqlQuery = "INSERT INTO user_transactions (id, user_id, company_stock_id, quantity, buy_price) VALUES (:id, :user_id, :company_stock_id, :quantity, :buy_price)"
	}
	_, err := database.DB.NamedExec(sqlQuery, userTransaction)
	if err != nil {
		return err
	}
	return nil
}
