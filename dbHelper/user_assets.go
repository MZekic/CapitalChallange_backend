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

func GetUserAssetsByTicker(ticker string, userID string) (models.UserAssets, error) {
	sqlQuery := "SELECT * FROM user_assets WHERE ticker = $1 AND user_id = $2"

	var userAssets models.UserAssets

	err := database.DB.Get(&userAssets, sqlQuery, ticker, userID)

	if err != nil {
		return userAssets, err
	}

	return userAssets, nil
}

func UpdateUserAssetsOnBuyByTicker(id string, quantity int) error {
	sqlQuery := "UPDATE user_assets SET quantity = $1 WHERE id = $2"

	_, err := database.DB.Exec(sqlQuery, quantity, id)

	if err != nil {
		return err
	}
	return nil
}

func UpdateUserAssetsOnUserTransaction(id string, quantity int) error {
	sqlQuery := "UPDATE user_assets SET quantity = $1 WHERE id = $2"

	_, err := database.DB.Exec(sqlQuery, quantity, id)

	if err != nil {
		return err
	}
	return nil
}
