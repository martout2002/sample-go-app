package threads

import (
	"encoding/json"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

func GetThreads(w http.ResponseWriter, r *http.Request) {
    rows, err := database.DB.Query("SELECT id, title, store_name, store_location, author_name, details, rating, comments FROM threads")
    if err != nil {
        http.Error(w, "Failed to fetch threads", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var threads []models.Thread
    for rows.Next() {
        var thread models.Thread
        if err := rows.Scan(&thread.ID, &thread.Title, &thread.StoreName, &thread.StoreLocation, &thread.AuthorName, &thread.Details, &thread.Rating, &thread.Comments); err != nil {
            http.Error(w, "Failed to scan thread", http.StatusInternalServerError)
            return
        }
        threads = append(threads, thread)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(threads)
}
