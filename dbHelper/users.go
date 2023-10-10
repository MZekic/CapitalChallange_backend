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

func CreateUserRecord(user models.Users) error {
	sqlQuery := "INSERT INTO users(id, username, email, password, current_game_number ) VALUES (:id, :username, :email, :password, :current_game_number)"
	_, err := database.DB.NamedExec(sqlQuery, user)
	if err != nil {
		return err
	}
	return nil
}

func CheckIfUsernameOrEmailExists(email string, username string) (int, int, error) {
	sqlQueryForEmail := "SELECT COUNT(*) from users WHERE email = $1"
	var emailExists int
	err := database.DB.Get(&emailExists, sqlQueryForEmail, email)
	if err != nil {
		return 0, 0, err
	}
	var usernameExists int
	sqlQueryForUsername := "SELECT COUNT(*) from users WHERE username = $1"
	err = database.DB.Get(&usernameExists, sqlQueryForUsername, username)
	if err != nil {
		return 0, 0, err
	}
	return emailExists, usernameExists, err
}

func GetUserByEmail(email string) (models.Users, error) {
	var user models.Users
	sqlQuery := "SELECT * FROM users WHERE email = $1"
	err := database.DB.Get(&user, sqlQuery, email)
	if err != nil {
		return user, err
	}

	return user, nil
}
