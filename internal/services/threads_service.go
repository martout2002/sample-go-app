package services

import (
	"errors"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

// CreateThread adds a new thread to the database
func CreateThread(thread models.Thread) error {
	_, err := database.DB.Exec(
		"INSERT INTO threads (title, store_name, store_location, author_name, details, rating, comments) VALUES (?, ?, ?, ?, ?, ?, ?)",
		thread.Title, thread.StoreName, thread.StoreLocation, thread.AuthorName, thread.Details, thread.Rating, thread.Comments,
	)
	return err
}

// DeleteThread removes a thread by ID
func DeleteThread(id string) error {
	result, err := database.DB.Exec("DELETE FROM threads WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("thread not found")
	}

	return nil
}

// GetThreads retrieves all threads from the database
func GetThreads() ([]models.Thread, error) {
	rows, err := database.DB.Query("SELECT id, title, store_name, store_location, author_name, details, rating, comments FROM threads")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var threads []models.Thread
	for rows.Next() {
		var thread models.Thread
		if err := rows.Scan(&thread.ID, &thread.Title, &thread.StoreName, &thread.StoreLocation, &thread.AuthorName, &thread.Details, &thread.Rating, &thread.Comments); err != nil {
			return nil, err
		}
		threads = append(threads, thread)
	}

	return threads, nil
}
