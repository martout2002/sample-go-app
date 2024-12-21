package routes

import (
	"net/http"

	"github.com/CVWO/sample-go-app/internal/handlers"
)

// SetupRouter registers all the routes
func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/threads", handlers.GetThreads)  // GET all threads
	mux.HandleFunc("/api/threads/add", handlers.CreateThread) // POST to create a thread

	return mux
}
