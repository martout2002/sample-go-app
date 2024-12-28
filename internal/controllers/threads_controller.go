package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/CVWO/sample-go-app/internal/models"
	"github.com/CVWO/sample-go-app/internal/services"
)

// CreateThread handles creating a new thread
func CreateThread(w http.ResponseWriter, r *http.Request) {
	var thread models.Thread
	if err := json.NewDecoder(r.Body).Decode(&thread); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Error decoding request payload: %v", err)
		return
	}

	err := services.CreateThread(thread)
	if err != nil {
		http.Error(w, "Failed to create thread", http.StatusInternalServerError)
		log.Printf("Error creating thread: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Thread created successfully"})
}

// DeleteThread handles deleting a thread
func DeleteThread(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		log.Println("Missing id parameter")
		return
	}

	err := services.DeleteThread(id)
	if err != nil {
		http.Error(w, "Failed to delete thread", http.StatusInternalServerError)
		log.Printf("Failed to delete thread with id %s: %v", id, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Thread deleted successfully"})
}

// GetThreads handles fetching all threads
func GetThreads(w http.ResponseWriter, r *http.Request) {
	threads, err := services.GetThreads()
	if err != nil {
		http.Error(w, "Failed to fetch threads", http.StatusInternalServerError)
		log.Printf("Error fetching threads: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(threads)
}

// LikeThread toggles the like status for a specific thread by a user
func LikeThread(w http.ResponseWriter, r *http.Request) {
	log.Println("LikeThread endpoint hit") // Log the endpoint hit

	// Get thread ID from query parameters
	threadIDStr := r.URL.Query().Get("id")
	if threadIDStr == "" {
		http.Error(w, "Thread ID is required", http.StatusBadRequest)
		log.Println("Missing thread ID in query parameters")
		return
	}

	// Convert thread ID to an integer
	threadID, err := strconv.Atoi(threadIDStr)
	if err != nil {
		http.Error(w, "Invalid thread ID", http.StatusBadRequest)
		log.Printf("Invalid thread ID: %s", threadIDStr)
		return
	}

	// Get the username from the header (simulating authentication)
	username := r.Header.Get("X-Username")
	if username == "" {
		http.Error(w, "Unauthorized: username missing", http.StatusUnauthorized)
		log.Println("Missing username in request header")
		return
	}

	// Use the username to get the user ID
	userID, err := services.GetUserIDByUsername(username)
	if err != nil {
		http.Error(w, "Failed to retrieve user ID", http.StatusInternalServerError)
		log.Printf("Error fetching user ID for username %s: %v", username, err)
		return
	}

	// Toggle the like status
	liked, err := services.ToggleLike(userID, threadID)
	if err != nil {
		http.Error(w, "Failed to toggle like status", http.StatusInternalServerError)
		log.Printf("Error toggling like status for thread ID %d by user ID %d: %v", threadID, userID, err)
		return
	}

	// Fetch the updated like count
	likeCount, err := services.GetLikesCount(threadID)
	if err != nil {
		http.Error(w, "Failed to fetch like count", http.StatusInternalServerError)
		log.Printf("Error fetching like count for thread ID %d: %v", threadID, err)
		return
	}

	// Respond with the updated like status and count
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"liked": liked,
		"likes": likeCount,
	})
}


// GetLikesCount handles fetching the like count for a specific thread
func GetLikesCount(w http.ResponseWriter, r *http.Request) {
	log.Println("GetLikesCount endpoint hit") // Log when the endpoint is accessed

	// Parse thread ID from query parameters
	threadIDStr := r.URL.Query().Get("id")
	if threadIDStr == "" {
		http.Error(w, "Thread ID is required", http.StatusBadRequest)
		log.Println("Thread ID missing from query parameters")
		return
	}

	// Convert thread ID to integer
	threadID, err := strconv.Atoi(threadIDStr)
	if err != nil {
		http.Error(w, "Invalid thread ID", http.StatusBadRequest)
		log.Printf("Invalid thread ID: %s", threadIDStr)
		return
	}

	// Fetch the like count from the service
	likeCount, err := services.GetLikesCount(threadID)
	if err != nil {
		http.Error(w, "Failed to fetch like count", http.StatusInternalServerError)
		log.Printf("Error fetching like count for thread ID %d: %v", threadID, err)
		return
	}

	// Respond with the like count
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"likes": likeCount})
}
