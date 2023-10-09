package dbHelper

import (
	database "capital-challenge-server/database"
	"capital-challenge-server/models"
)

func InsertUserTransaction(userTransaction models.UserTransactions) error {
	sqlQuery := "INSERT INTO user_transactions (id, user_id, company_stock_id, buy_or_sell, quantity, game_number) VALUES (:id, :user_id, :company_stock_id, :buy_or_sell, :quantity, :game_number)"
	_, err := database.DB.NamedExec(sqlQuery, userTransaction)
	if err != nil {
		return err
	}
	return nil
}
