package routes

import (
	"net/http"

	"github.com/CVWO/sample-go-app/internal/handlers/threads"
	"github.com/CVWO/sample-go-app/internal/handlers/users"
)

func SetupRouter() *http.ServeMux {
    mux := http.NewServeMux()

    mux.HandleFunc("/api/threads", threads.GetThreads)  // GET all threads
    mux.HandleFunc("/api/threads/add", threads.CreateThread) // POST to create a thread
    mux.HandleFunc("/api/threads/delete", threads.DeleteThread) // DELETE a thread

    // User-related routes
    mux.HandleFunc("/api/users", users.HandleList)        // GET all users
    mux.HandleFunc("/api/login", users.HandleLogin)       // POST to login
    mux.HandleFunc("/api/signup", users.SignupUser)       // POST to signup

    return mux
}
