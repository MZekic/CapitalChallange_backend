package dbHelper

import (
	database "capital-challenge-server/database"
	"capital-challenge-server/models"
)

func InsertUserTransactionOnBuy(userTransaction models.UserTransactions) error {
	sqlQuery := "INSERT INTO user_transactions (id, user_id, company_stock_id, quantity, buy_price) VALUES (:id, :user_id, :company_stock_id, :quantity, :buy_price)"
	_, err := database.DB.NamedExec(sqlQuery, userTransaction)
	if err != nil {
		return err
	}
	return nil
}

func InsertUserTransactionOnSell(userTransaction models.UserTransactions) error {
	sqlQuery := "INSERT INTO user_transactions (id, user_id, company_stock_id, quantity, sell_price) VALUES (:id, :user_id, :company_stock_id, :quantity, :sell_price)"
	_, err := database.DB.NamedExec(sqlQuery, userTransaction)
	if err != nil {
		return err
	}
	return nil
}
