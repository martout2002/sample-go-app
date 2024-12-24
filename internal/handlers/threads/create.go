package threads

import (
	"encoding/json"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

func CreateThread(w http.ResponseWriter, r *http.Request) {
    var thread models.Thread
    if err := json.NewDecoder(r.Body).Decode(&thread); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    _, err := database.DB.Exec(
        "INSERT INTO threads (title, store_name, store_location, author_name, details, rating, comments) VALUES (?, ?, ?, ?, ?, ?, ?)",
        thread.Title, thread.StoreName, thread.StoreLocation, thread.AuthorName, thread.Details, thread.Rating, thread.Comments,
    )
    if err != nil {
        http.Error(w, "Failed to create thread", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}
