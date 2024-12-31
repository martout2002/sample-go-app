package services

import (
	"encoding/json"
	"errors"

	"searchFoodBackend/internal/database"
	"searchFoodBackend/internal/models"
)

// CreateThread adds a new thread to the database
func CreateThread(thread models.Thread) error {
	_, err := database.DB.Exec(
		"INSERT INTO threads (title, store_name, store_location, author_name, details, rating, comments, likes) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		thread.Title, thread.StoreName, thread.StoreLocation, thread.AuthorName, thread.Details, thread.Rating, thread.Comments, "[]",
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

// ToggleLike toggles the like status for a thread by a user
func ToggleLike(userID, threadID int) (bool, error) {
	// Retrieve the thread
	thread, err := GetThreadByID(threadID)
	if err != nil {
		return false, err
	}

	username, err := GetUsernameByID(userID)
	if err != nil {
		return false, errors.New("failed to retrieve username")
	}

	// Check if the user already liked the thread
	for i, user := range thread.Likes {
		if user == username {
			// Remove the like
			thread.Likes = append(thread.Likes[:i], thread.Likes[i+1:]...)
			err = UpdateThreadLikes(threadID, thread.Likes)
			if err != nil {
				return false, err
			}
			return false, nil
		}
	}

	// Add the like
	thread.Likes = append(thread.Likes, username)
	err = UpdateThreadLikes(threadID, thread.Likes)
	if err != nil {
		return false, err
	}
	return true, nil
}

// UpdateThreadLikes updates the likes for a specific thread in the database
func UpdateThreadLikes(threadID int, likes []string) error {
	likesJSON, err := json.Marshal(likes)
	if err != nil {
		return err
	}

	_, err = database.DB.Exec("UPDATE threads SET likes = ? WHERE id = ?", string(likesJSON), threadID)
	return err
}

// GetLikesCount retrieves the total number of likes for a specific thread
func GetLikesCount(threadID int) (int, error) {
	thread, err := GetThreadByID(threadID)
	if err != nil {
		return 0, err
	}

	return len(thread.Likes), nil
}

func GetThreadByID(threadID int) (*models.Thread, error) {
	var thread models.Thread
	var likesJSON string

	// Fetch thread details, including likes
	err := database.DB.QueryRow(
		"SELECT id, title, store_name, store_location, author_name, details, rating, comments, likes FROM threads WHERE id = ?",
		threadID,
	).Scan(&thread.ID, &thread.Title, &thread.StoreName, &thread.StoreLocation, &thread.AuthorName, &thread.Details, &thread.Rating, &thread.Comments, &likesJSON)
	if err != nil {
		return nil, err
	}

	// Deserialize likes
	var likes []string
	if err := json.Unmarshal([]byte(likesJSON), &likes); err != nil {
		return nil, err
	}

	thread.Likes = likes // Assign the array of usernames

	return &thread, nil
}

func GetThreads() ([]models.Thread, error) {
	rows, err := database.DB.Query("SELECT id, title, store_name, store_location, author_name, details, rating, comments, likes FROM threads")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var threads []models.Thread
	for rows.Next() {
		var thread models.Thread
		var likesJSON string

		// Fetch thread data
		if err := rows.Scan(&thread.ID, &thread.Title, &thread.StoreName, &thread.StoreLocation, &thread.AuthorName, &thread.Details, &thread.Rating, &thread.Comments, &likesJSON); err != nil {
			return nil, err
		}

		// Deserialize likes into a []string
		var likes []string
		if err := json.Unmarshal([]byte(likesJSON), &likes); err != nil {
			return nil, err
		}

		thread.Likes = likes // Assign the array of usernames
		threads = append(threads, thread)
	}

	return threads, nil
}
