package users

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

// Get all threads
func GetThreads(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling GET /api/threads request") // Log when the endpoint is hit

	// Check if database connection is initialized
	if database.DB == nil {
		log.Println("Database connection is nil")
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}

	// Query the database
	rows, err := database.DB.Query("SELECT id, title, store_name, store_location, author_name, details, rating, comments FROM threads")
	if err != nil {
		log.Printf("Error querying database: %v", err)
		http.Error(w, "Failed to fetch threads", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var threads []models.Thread
	for rows.Next() {
		var thread models.Thread
		if err := rows.Scan(&thread.ID, &thread.Title, &thread.StoreName, &thread.StoreLocation, &thread.AuthorName, &thread.Details, &thread.Rating, &thread.Comments); err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Failed to scan thread", http.StatusInternalServerError)
			return
		}
		threads = append(threads, thread)
	}

	// Log the fetched data
	log.Printf("Fetched threads: %+v", threads)

	// Respond with JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(threads); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}


func CreateThread(w http.ResponseWriter, r *http.Request) {
	log.Println("Received POST /api/threads/add") // Log the endpoint hit

	var thread models.Thread
	if err := json.NewDecoder(r.Body).Decode(&thread); err != nil {
		log.Println("Error decoding request body:", err) // Log decoding errors
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	log.Printf("Decoded thread: %+v\n", thread) // Log the decoded thread payload

	// Insert into database
	query := "INSERT INTO threads (title, store_name, store_location, author_name, details, rating, comments) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := database.DB.Exec(query, thread.Title, thread.StoreName, thread.StoreLocation, thread.AuthorName, thread.Details, thread.Rating, thread.Comments)
	if err != nil {
		log.Println("Error inserting into database:", err) // Log database errors
		http.Error(w, "Failed to create thread", http.StatusInternalServerError)
		return
	}

	log.Println("Thread created successfully") // Log success
	w.WriteHeader(http.StatusCreated)
}

func DeleteThread(w http.ResponseWriter, r *http.Request) {
    // Ensure the method is DELETE
    if r.Method != http.MethodDelete {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Parse the thread ID from query parameters
    idStr := r.URL.Query().Get("id")
    if idStr == "" {
        http.Error(w, "Thread ID is required", http.StatusBadRequest)
        return
    }

    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid thread ID", http.StatusBadRequest)
        return
    }

    // Delete the thread from the database
    db := database.DB
    _, err = db.Exec("DELETE FROM threads WHERE id = ?", id)
    if err != nil {
        http.Error(w, "Failed to delete thread", http.StatusInternalServerError)
        return
    }

    // Respond with success
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Thread deleted successfully"))
}

