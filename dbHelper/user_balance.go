package dbHelper

import (
	database "capital-challenge-server/database"
	"capital-challenge-server/models"
)

func GetUserBalanceByUserID(userID string) (models.UserBalance, error) {
	var userBalance models.UserBalance
	sqlQuery := "SELECT * FROM user_balance WHERE user_id = $1"
	err := database.DB.Get(&userBalance, sqlQuery, userID)
	if err != nil {
		return userBalance, err
	}

	return userBalance, nil
}

func UpdateUserBalance(userID string, updatedBalance float32) error {
	sqlQuery := "UPDATE user_balance SET current_balance = $1 WHERE user_id= $2"
	_, err := database.DB.Exec(sqlQuery, updatedBalance, userID)

	if err != nil {
		return err
	}

	return nil
}

func CreateUserBalance(userBalance models.UserBalance) error {
	sqlQuery := "INSERT INTO user_balance(id, user_id, starting_balance, current_balance) VALUES (:id, :user_id, :starting_balance, :current_balance)"
	_, err := database.DB.NamedExec(sqlQuery, userBalance)
	if err != nil {
		return err
	}
	return nil
}
