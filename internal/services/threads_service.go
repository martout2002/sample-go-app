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

// ToggleLike toggles the like status for a thread
func ToggleLike(userID, threadID int) (bool, error) {
	// Check if the user already liked the thread
	var exists bool
	err := database.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM user_likes WHERE user_id = ? AND thread_id = ?)",
		userID, threadID,
	).Scan(&exists)
	if err != nil {
		return false, err
	}

	if exists {
		// Unlike: Remove the like
		_, err = database.DB.Exec("DELETE FROM user_likes WHERE user_id = ? AND thread_id = ?", userID, threadID)
		if err != nil {
			return false, err
		}
		return false, nil // False means unliked
	}

	// Like: Add a new like
	_, err = database.DB.Exec("INSERT INTO user_likes (user_id, thread_id) VALUES (?, ?)", userID, threadID)
	if err != nil {
		return false, err
	}
	return true, nil // True means liked
}

// GetLikesCount retrieves the number of likes for a specific thread
func GetLikesCount(threadID int) (int, error) {
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM user_likes WHERE thread_id = ?", threadID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

