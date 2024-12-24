package controllers

import (
	"encoding/json"
	"log"
	"net/http"

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
