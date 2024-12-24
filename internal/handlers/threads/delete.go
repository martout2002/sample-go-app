package threads

import (
	"log"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/database"
)

func DeleteThread(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "Missing id parameter", http.StatusBadRequest)
        return
    }

    _, err := database.DB.Exec("DELETE FROM threads WHERE id = ?", id)
    if err != nil {
        log.Printf("Failed to delete thread with id %s: %v", id, err)
        http.Error(w, "Failed to delete thread", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Thread deleted successfully"))
}

