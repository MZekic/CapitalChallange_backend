package dbHelper

import (
	database "capital-challenge-server/database"
	"capital-challenge-server/models"
)

func InsertUserAsset(userAsset models.UserAssets) error {
	sqlQuery := "INSERT INTO user_assets (id, user_id, ticker, quantity, game_number) VALUES (:id, :user_id, :ticker, :quantity, :game_number)"
	_, err := database.DB.NamedExec(sqlQuery, userAsset)
	if err != nil {
		return err
	}
	return nil
}
