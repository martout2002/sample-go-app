package services

import (
	"errors"

	"github.com/CVWO/sample-go-app/internal/database"
)

// AddLike adds a like for a thread by a specific user
func AddLike(userID, threadID int) error {
	// Check if the user already liked the thread
	var exists bool
	err := database.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM user_likes WHERE user_id = ? AND thread_id = ?)",
		userID, threadID,
	).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user has already liked this thread")
	}

	// Insert the like into the table
	_, err = database.DB.Exec(
		"INSERT INTO user_likes (user_id, thread_id) VALUES (?, ?)",
		userID, threadID,
	)
	return err
}

// RemoveLike removes a like for a thread by a specific user
func RemoveLike(userID, threadID int) error {
	// Delete the like from the table
	_, err := database.DB.Exec(
		"DELETE FROM user_likes WHERE user_id = ? AND thread_id = ?",
		userID, threadID,
	)
	return err
}

// HasUserLiked checks if a user has liked a specific thread
func HasUserLiked(userID, threadID int) (bool, error) {
	var exists bool
	err := database.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM user_likes WHERE user_id = ? AND thread_id = ?)",
		userID, threadID,
	).Scan(&exists)
	return exists, err
}