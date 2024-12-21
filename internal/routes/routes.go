package routes

import (
	"net/http"

	"github.com/CVWO/sample-go-app/internal/handlers/users" // Update to correct package
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Register the /api/threads endpoint
	mux.HandleFunc("/api/threads", users.GetThreads)
	mux.HandleFunc("/api/threads/add", users.CreateThread)  // Registers the /api/threads/add route

	return mux
}
