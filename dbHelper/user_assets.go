package dbHelper

import (
	database "capital-challenge-server/database"
	"capital-challenge-server/models"
)

func InsertUserAsset(userAsset models.UserAssets) error {
	sqlQuery := "INSERT INTO user_assets (id, user_id, ticker, quantity, company_stock_id, price_per_unit) VALUES (:id, :user_id, :ticker, :quantity, :company_stock_id, :price_per_unit)"
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

func GetUserAssetsByCompanyStock(ticker string, userID string) (models.UserAssets, error) {
	sqlQuery := "SELECT * FROM user_assets WHERE company_stock_id = $1 AND user_id = $2"

	var userAssets models.UserAssets

	err := database.DB.Get(&userAssets, sqlQuery, ticker, userID)

	if err != nil {
		return userAssets, err
	}

	return userAssets, nil
}

func UpdateUserAssets(id string, quantity int) error {
	sqlQuery := "UPDATE user_assets SET quantity = $1 WHERE id = $2"

	_, err := database.DB.Exec(sqlQuery, quantity, id)

	if err != nil {
		return err
	}
	return nil
}

func DeleteUserAssets(id string) error {
	sqlQuery := "DELETE FROM user_assets WHERE id = $1"

	_, err := database.DB.Exec(sqlQuery, id)

	if err != nil {
		return err
	}
	return nil
}

func GetUserAssets(userID string) ([]models.UserAssets, error) {
	sqlQuery := "SELECT * FROM user_assets WHERE user_id = $1"

	var userAssets []models.UserAssets

	err := database.DB.Select(&userAssets, sqlQuery, userID)

	if err != nil {
		return userAssets, err
	}

	return userAssets, nil
}

type UserAssetsQuantity struct {
	Ticker   string `json:"ticker" db:"ticker"`
	Quantity int    `json:"quantity" db:"sum"`
}

func GetUserAssetsQuantityByTicker(userID string) ([]UserAssetsQuantity, error) {
	sqlQuery := "SELECT ticker, SUM(quantity) FROM user_assets WHERE user_id = $1 GROUP BY ticker, quantity"

	var userAssetsQuantity []UserAssetsQuantity

	err := database.DB.Select(&userAssetsQuantity, sqlQuery, userID)

	if err != nil {
		return userAssetsQuantity, err
	}

	return userAssetsQuantity, nil
}
