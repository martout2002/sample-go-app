package services

import (
	"database/sql"
	"errors"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

func GetAllUsers() ([]models.User, error) {
	rows, err := database.DB.Query("SELECT id, username FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func CreateUser(user models.User) error {
	// Check if username exists
	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", user.Username).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("username already exists")
	}

	_, err = database.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	return err
}

func VerifyUser(loginReq models.LoginRequest) (bool, error) {
	var storedPassword string
	err := database.DB.QueryRow("SELECT password FROM users WHERE username = ?", loginReq.Username).Scan(&storedPassword)
	if err == sql.ErrNoRows {
		return false, nil // User not found
	}
	if err != nil {
		return false, err // Database error
	}

	// Return true if passwords match
	return storedPassword == loginReq.Password, nil
}

