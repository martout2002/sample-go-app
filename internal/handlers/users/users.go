package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

func HandleList(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling list users request...")

	// Check if the database connection is initialized
	db := database.DB
	if db == nil {
		http.Error(w, "Database connection not initialized", http.StatusInternalServerError)
		log.Println("Database connection not initialized")
		return
	}

	// Fetch users from the database
	rows, err := db.Query("SELECT id, username FROM users")
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		log.Println("Error retrieving users:", err)
		return
	}
	defer rows.Close()

	// Parse the results
	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			http.Error(w, "Failed to scan user data", http.StatusInternalServerError)
			log.Println("Error scanning user data:", err)
			return
		}
		users = append(users, user)
	}

	// Convert to JSON and send response
	response, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
		log.Println("Error encoding users:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	log.Println("Successfully retrieved users")
}
