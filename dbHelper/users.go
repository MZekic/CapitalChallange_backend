package dbHelper

import (
	database "capital-challenge-server/database"
	"capital-challenge-server/models"
	"database/sql"
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
	sqlQuery := "INSERT INTO users(id, username, email, password) VALUES (:id, :username, :email, :password)"
	_, err := database.DB.NamedExec(sqlQuery, user)
	if err != nil {
		return err
	}
	return nil
}

func CheckIfUsernameOrEmailExists(email string, username string) (bool, error) {
	sqlQuery := "SELECT * from users WHERE email = $1 OR username = $2"
	var ifExists models.Users
	err := database.DB.Get(&ifExists, sqlQuery, email, username)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	} else if err == sql.ErrNoRows {
		return false, nil
	}

	return true, nil
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
