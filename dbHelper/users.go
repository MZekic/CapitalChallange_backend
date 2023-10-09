package dbHelper

import (
	database "capital-challenge-server/database"
	"capital-challenge-server/models"
)

func GetUserByID(userID string) (models.Users, error) {

	var user models.Users
	sqlQuery := "SELECT * FROM users WHERE id = $1"
	err := database.DB.Get(&user, sqlQuery, userID)
	if err != nil {
		return user, err
	}

	return user, nil
}
